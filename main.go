package main

import (
	"miniorder/common"
	"miniorder/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	db := common.InitDB()
	d, _ := db.DB()
	defer d.Close()

	e = routes.CollectRouter(e)
	panic(e.Run())
}
