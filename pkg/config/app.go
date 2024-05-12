package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	d, err := gorm.Open(mysql.Open("root:root@(localhost)/test?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	DB = d
}

func GetDB() *gorm.DB {
	if DB == nil {
		Connect()
	}
	return DB
}
