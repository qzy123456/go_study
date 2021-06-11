package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

var dbb *gorm.DB
var err error

type Userrr struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	dbb, err = gorm.Open("sqlite3", "./api.db")
	//db, err = gorm.Open("sqlite3", "./api.db")
	//使用mysql, gorm.Open(“mysql”, “user:pwd@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local”)
	if err != nil {
		log.Fatal("db connect error")
	}
	defer dbb.Close() //延时调用函数
	dbb.AutoMigrate(&Userrr{})

	r := gin.Default()
	r.Use(Cors())
	r.GET("/users", indexs)         //获取所有用户
	r.GET("/users/:id", show)       //根据id获取用户
	r.POST("/users", store)         //保存新用户
	r.PUT("/users/:id", update)     //根据id更新用户
	r.DELETE("/users/:id", destroy) //根据id删除用户
	_ = r.Run(":7777")
}

func indexs(c *gin.Context) {
	var users []Userrr
	dbb.Find(&users)
	c.JSON(200, gin.H{"code": 1, "message": "success", "data": users})
}

func show(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Userrr
	dbb.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, user)
}

func store(c *gin.Context) {
	var user Userrr
	_ = c.BindJSON(&user) //绑定一个请求主体到一个类型
	dbb.Create(&user)
	c.JSON(200, user)
}

func update(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Userrr
	dbb.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		dbb.Save(&user)
		c.JSON(200, user)
	}
}

func destroy(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Userrr
	dbb.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		dbb.Delete(&user)
		c.JSON(200, gin.H{"message": "delete success"})
	}
}
