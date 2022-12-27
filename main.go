package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20); not null"`
	Telephone string `gorm:"varchar(11); not null"`
	Password  string `gorm:"size:255; not null"`
}

func InitDB() *gorm.DB {

	dsn := "root:123456@tcp(127.0.0.1:3306)/ginessential?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("failed to connect database, err:" + err.Error())
	}
	return db
}

func main() {
	e := gin.Default()
	db := InitDB()
	d, _ := db.DB()
	defer d.Close()

	e.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号码必须为11位",
			})
			return
		}
		if len(password) < 5 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不能少于6位",
			})
			return
		}
		// 如果名称没有传，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, telephone, password)

		//判断手机号码是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户已存在",
			})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	panic(e.Run())
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
