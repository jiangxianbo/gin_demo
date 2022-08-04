package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建一个默认的路由引擎
	// 默认使用了2个中间件 Logger(), Recovery()
	r := gin.Default()

	//// 也可以创建不带中间件的路由
	//// r := gin.New()
	//// 2.绑定路由，执行函数
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello world!")
	//	// c.JSON：返回JSON格式的数据
	//	// c.JSON(200, gin.H{
	//	//	 "message": "Hello world!",
	//	// })
	//})
	//r.POST("/xxxPost", getting)

	// api参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})

	// url参数
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	// 3.监听端口，默认8080
	r.Run(":8000")
}

func getting(c *gin.Context) {

}
