package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:1234/efine_03?charset=utf8&parseTime=True&Loc=Local")

	if err != nil {
		panic(err.Error())
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
