package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	dsn := "jan:123456@tcp(192.168.233.128:3306)/ginessential?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("failed to connect database, err:" + err.Error())
	}
	// db.AutoMigrate(&User{})  创建user表
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
