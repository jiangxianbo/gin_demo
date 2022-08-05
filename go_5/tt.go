package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	c.String(http.StatusOK, name+" is "+action)
	//})

	// url参数
	//r.GET("/welcome", func(c *gin.Context) {
	//	name := c.DefaultQuery("name", "Jack")
	//	c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	//})

	// 表单参数
	//r.POST("/form", func(c *gin.Context) {
	//	// 表单参数设置默认值
	//	type1 := c.DefaultPostForm("type", "alert")
	//	// 接受其他的
	//	username := c.PostForm("username")
	//	password := c.PostForm("password")
	//	hobbys := c.PostFormArray("hobby")
	//	c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v", type1, username, password, hobbys))
	//})

	// 上传文件
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "./go_5/page/"+file.Filename)
		c.String(200, fmt.Sprintf("%s upload!", file.Filename))
	})

	// 3.监听端口，默认8080
	r.Run(":8000")
}

func getting(c *gin.Context) {

}
