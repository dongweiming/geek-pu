package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/dongweiming/geek-pu/errors"
	. "github.com/dongweiming/geek-pu/models"

	"github.com/gin-gonic/gin"
)

const (
	APPID       string = "wxf0855fcd9689fd67"
	SECRET      string = "c9389e5cb38cb2e7e4e4c939367d9479"
	ACCESSTOKEN string = "access-token"
)

var (
	areaMap = map[string]string{
		"au": "澳版",
		"us": "美版",
	}
	c = cache.New(60*time.Minute, 10*time.Minute)
)

func HttpRequest(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return body, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}

func WeChatLogin(code string) ([]byte, error) {
	base_url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v"+
		"&secret=%v&js_code=%v&grant_type=authorization_code",
		APPID, SECRET, code)
	body, err := HttpRequest(base_url)
	return body, err
}

func GetAccessToken() (string, error) {
	if x, found := c.Get(ACCESSTOKEN); found {
		return x.(string), nil
	}
	base_url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		APPID, SECRET)
	body, err := HttpRequest(base_url)
	var resp GetAccessTokenResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}
	if resp.Errcode != 0 {
		return "", GeekError{resp.Errmsg}
	}
	c.Set(ACCESSTOKEN, resp.AccessToken, 60*time.Minute)
	return resp.AccessToken, nil
}

func SubscribeMessage(openid string) {
	token, err := GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", token)
	values := &SubscribeReq{
		AccessToken: token,
		Touser:      openid,
		TemplateId:  624,
		Data: SubscribeData{
			Thing2:  WxValue{Value: "测试"},
			Phrase5: WxValue{Value: "成功"},
		},
	}

	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	defer resp.Body.Close()
}

func ResponseJsonError(c *gin.Context, code int) {
	msg := errors.GetMsg(code)
	c.JSON(code, gin.H{"message": msg})
}

func PostCode2Session(c *gin.Context, code string) string {
	var resp WeChatAuthResp

	res, err := WeChatLogin(code)
	if err != nil {
		ResponseJsonError(c, errors.INVALID_PARAMS)
		return ""
	}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		ResponseJsonError(c, errors.INVALID_PARAMS)
		return ""
	}

	if resp.ErrCode != 0 {
		ResponseJsonError(c, errors.INVALID_PARAMS)
		return ""
	}
	return resp.OpenId
}

func Subscribe(c *gin.Context) {
	gid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"ok": 0,
		})
		return
	}
	openId := c.Request.Header.Get("x-corran-token")
	if openId == "" {
		c.JSON(200, gin.H{
			"ok": 0,
		})
		return
	}
	sub := Subscription{Uid: openId, Gid: gid}
	db := GetDB()
	defer db.Close()
	if db.Create(&sub).Error != nil {
		c.JSON(200, gin.H{
			"ok": 0,
		})
	} else {
		SubscribeMessage(openId)
		c.JSON(200, gin.H{
			"ok": 1,
		})
	}
}

func UnSubscribe(c *gin.Context) {
	gid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"ok": 0,
		})
		return
	}
	openId := c.Request.Header.Get("x-corran-token")
	if openId == "" {
		c.JSON(200, gin.H{
			"ok": 0,
		})
		return
	}
	db := GetDB()
	defer db.Close()
	db.Where("id = ? AND uid = ?", gid, openId).Delete(&Subscription{})
	c.JSON(200, gin.H{
		"ok": 1,
	})
}

func GetGameList(c *gin.Context) {
	q := c.Query("q")
	filter := c.Query("filter")

	openId := c.Request.Header.Get("x-corran-token")
	if openId == "" {
		c.JSON(400, gin.H{"error": "Wrong request"})
		return
	}

	var games []Game
	var subs []Subscription
	var sub Subscription
	db := GetDB()
	defer db.Close()

	if q != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", q))
	}

	if filter != "" {
		if strings.HasPrefix(filter, "zone") {
			areas := strings.Split(filter, ":")
			if area, ok := areaMap[areas[1]]; ok {
				db = db.Where("area = ?", area)
			}
		} else if strings.HasPrefix(filter, "order") {
			orders := strings.Split(filter, ":")
			order := orders[1]
			if strings.Contains(order, "rating") {
				db = db.Order("rating desc")
			} else if strings.Contains(order, "price") {
				sorts := strings.Split(order, "-")
				db = db.Order(fmt.Sprintf("price %s", sorts[1]))
			}
		} else if strings.HasPrefix(filter, "edition") {
			db = db.Where("id >= ? AND id <= ?", 13, 15).Order("release_date")
		} else if strings.HasPrefix(filter, "mine:subscription") {
			db.Select("gid").Where("uid = ?", openId).Find(&subs)
			if len(subs) == 0 {
				c.JSON(200, gin.H{
					"games": games,
				})
				return
			}
			var ids []int
			for _, sub = range subs {
				ids = append(ids, sub.Gid)
			}
			db = db.Where(ids)
		} else if strings.HasPrefix(filter, "in-stock") {
			db = db.Where("quantity > 0")
		}
	}

	db.Find(&games)

	if len(subs) == 0 {
		db.Select("gid").Where("uid = ?", openId).Find(&subs)
	}
	if len(subs) != 0 {
		var ids []int
		for _, sub = range subs {
			ids = append(ids, sub.Gid)
		}

		for index := range games {
			game := games[index]
			if Contains(ids, game.ID) {
				game.Subscribed = true
				games[index] = game
			}
		}
	}

	c.JSON(200, gin.H{
		"games": games,
	})

}

func main() {
	r := gin.Default()
	r.POST("/api/login", func(c *gin.Context) {
		var code Code
		c.BindJSON(&code)

		openId := PostCode2Session(c, code.Data)
		c.JSON(200, gin.H{
			"openid": openId,
		})

	})
	r.GET("/api/games", GetGameList)
	r.POST("/api/game/:id/subscribe", Subscribe)
	r.POST("/api/game/:id/unsubscribe", UnSubscribe)

	r.Run(":80")
}
