package common

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	// dsn := "jan:123456@tcp(192.168.233.128:3306)/ginessential?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: driverName,
		DSN:        dsn,
	}), &gorm.Config{})
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
