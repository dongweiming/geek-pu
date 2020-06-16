package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/dongweiming/geek-pu/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
)

func initdb() {
	db := GetDB()
	defer db.Close()
	for _, table := range []interface{}{Game{}, Subscription{}} {
		db.DropTableIfExists(table)
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(table)
	}
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "initdb",
				Aliases: []string{"i"},
				Usage:   "Initialize the database",
				Action: func(c *cli.Context) error {
					initdb()
					fmt.Println(Bold(Cyan("Done!")))
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add game",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "title",
						Aliases:  []string{"t"},
						Usage:    "Game Title",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "cover",
						Aliases:  []string{"c"},
						Usage:    "Cover file name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "release_date",
						Aliases:  []string{"r"},
						Usage:    "Release date",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "rating",
						Aliases:  []string{"s"},
						Usage:    "Douban rating score",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "douban_id",
						Aliases:  []string{"d"},
						Usage:    "Douban id",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "area",
						Aliases:     []string{"a"},
						Usage:       "Area verson",
						Value:       "Switch",
						DefaultText: "Switch",
					},
					&cli.StringFlag{
						Name:     "languages",
						Aliases:  []string{"l"},
						Usage:    "Suuported language list",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "platforms",
						Aliases:  []string{"p"},
						Usage:    "Suuported platforms",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					title := c.String("title")
					cover := c.String("cover")
					release_date := c.String("release_date")
					rating := c.Float64("rating")
					douban_id := c.Int("douban_id")
					area := c.String("area")
					languages := c.String("languages")
					platforms := c.String("platforms")
					price := c.Float64("price")
					quantity := c.Int("quantity")
					game := Game{Title: title, Cover: cover, ReleaseDate: release_date, Rating: rating,
						Area: area, Languages: languages, Platforms: platforms, DoubanID: douban_id,
						Price: price, Quantity: quantity}
					db := GetDB()
					defer db.Close()
					if db.NewRecord(game) {
						db.Create(&game)
						fmt.Println(Bold(Cyan("Done!")))
					}
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update price & quantity",
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:     "price",
						Aliases:  []string{"p"},
						Usage:    "Price",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "id",
						Aliases:  []string{"i"},
						Usage:    "Game id",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "quantity",
						Aliases:  []string{"t"},
						Usage:    "Quantity",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					id := c.Int("id")
					price := c.Float64("price")
					quantity := c.Int("quantity")
					var game Game
					db := GetDB()
					if db.Where(id).First(&game).RecordNotFound() {
						output := fmt.Sprintf("Game(id=%d) is not exists", id)
						fmt.Println(Bold(Red(output)))
						return nil
					}
					game.Price = price
					game.Quantity = quantity
					db.Save(&game)
					output := fmt.Sprintf("Updated: ID(%d) -> Price(%f), Quantity(%d)", id, price, quantity)
					fmt.Println(Bold(Green(output)))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
