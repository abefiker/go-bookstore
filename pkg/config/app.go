package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var (
	db *gorm.DB
)

func Connect() {
	dbUser := "root"
	dbPass := "MySql1994@be"


	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass)


	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
