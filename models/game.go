package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	uri := GetDbUri()
	db, err := gorm.Open("mysql", uri)
	// defer db.Close()
	checkError(err)
	return db
}

type Game struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"_"`
	UpdatedAt   time.Time `json:"_"`
	DeletedAt   time.Time `json:"_"`
	Title       string    `gorm:"type:varchar(100);unique_index" json:"title"`
	Cover       string    `gorm:"type:varchar(200)" json:"cover"`
	ReleaseDate string    `gorm:"type:varchar(200)" json:"release_date"` // 原谅我
	Rating      float64   `gorm:"type:decimal(10, 2)" json:"rating"`
	Area        string    `gorm:"type:varchar(100)" json:"area"`
	Languages   string    `gorm:"type:varchar(100)" json:"languages"`
	Platforms   string    `gorm:"type:varchar(100)" json:"platforms"`
	DoubanID    int       `gorm:"type:int" json:"douban_id"`
	Price       float64   `gorm:"type:decimal(10, 2)" json:"price"`
	Quantity    int       `gorm:"type:int" json:"quantity"`
	Desc        string    `gorm:"type:varchar(200)" json:"desc"`
	Subscribed  bool      `sql:"-" json:"subscribed"`
}

func (game Game) IsRefresh() bool {
	shiftTime := time.Now().AddDate(0, 0, -180)
	t, _ := time.Parse("2006-01-02", game.ReleaseDate)
	return t.After(shiftTime)
}

func (game *Game) MarshalJSON() ([]byte, error) {
	type Alias Game
	return json.Marshal(&struct {
		*Alias
		Refresh bool `json:"refresh"`
	}{
		Alias:   (*Alias)(game),
		Refresh: game.IsRefresh(),
	})
}

type Subscription struct {
	gorm.Model
	Uid string `gorm:"type:varchar(100)" json:"uid"`
	Gid int    `gorm:"type:int" json:"gid"`
}
