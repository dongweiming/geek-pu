package models

type WeChatAuthResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type Code struct {
	Data string `json:"code" binding:"required"`
}

type WxValue struct {
	Value string `json:"value"`
}

type SubscribeData struct {
	Thing2  WxValue `json:"thing2"`
	Phrase5 WxValue `json:"phrase5"`
}

type SubscribeReq struct {
	AccessToken string        `json:"access_token"`
	Touser      string        `json:"touser"`
	TemplateId  int           `json:"template_id"`
	Data        SubscribeData `json:"data"`
}

type GetAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}
