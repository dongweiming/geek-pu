package models

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
)

const DefaultFilePath = "config.toml"

func ReadConfig() *toml.Tree {
	path, _ := filepath.Abs(DefaultFilePath)
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
