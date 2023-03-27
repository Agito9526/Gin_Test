package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconnect *gorm.DB

var err error

func DD() {
	dsn := "root:9526@(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local"
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
