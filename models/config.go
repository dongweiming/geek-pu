package models

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/pelletier/go-toml"
)

const DefaultFilePath = "../config.toml"

func ReadConfig() *toml.Tree {
	_, filename, _, _ := runtime.Caller(1)
	path := path.Join(path.Dir(filename), "../config.toml")
	var config *toml.Tree
	if _, err := os.Stat(path); os.IsNotExist(err) {
		config, _ = toml.Load(`
[db]
user = "root"
password = ""
host = "localhost"
port = "3306"
database = "test"
charset = "utf8"

[product]
edition_ids = [13, 14]
`)
	} else {
		config, _ = toml.LoadFile(path)
	}
	return config
}

func GetDbUri() string {
	config := ReadConfig()
	dbConfig := config.Get("db").(*toml.Tree)
	user := dbConfig.Get("user").(string)
	passwd := dbConfig.Get("password").(string)
	host := dbConfig.Get("host").(string)
	port := dbConfig.Get("port").(string)
	database := dbConfig.Get("database").(string)
	charset := dbConfig.Get("charset").(string)

	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&charset=%s", user, passwd, host, port, database, charset)
}

func GetEditionIds() []int64 {
	config := ReadConfig()
	ids := config.GetArray("product.edition_ids").([]int64)
	return ids
}
