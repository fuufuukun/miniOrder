package main

import (
	"miniorder/common"
	"miniorder/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()

	db := common.InitDB()
	d, _ := db.DB()
	defer d.Close()

	e := gin.Default()
	e = routes.CollectRouter(e)
	port := viper.GetString("server.port")
	if port != "" {
		panic(e.Run(":" + port))
	}
	panic(e.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
