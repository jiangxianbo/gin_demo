package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user"  xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password"  xml:"password" binding:"required"`
}

func main() {
	// 1.创建一个默认的路由引擎
	// 默认使用了2个中间件 Logger(), Recovery()
	r := gin.Default()

	// 也可以创建不带中间件的路由
	// r := gin.New()

	// 2.绑定路由，执行函数
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello world!")
	//	// c.JSON：返回JSON格式的数据
	//	// c.JSON(200, gin.H{
	//	//	 "message": "Hello world!",
	//	// })
	//})
	//r.POST("/xxxPost", getting)

	// 3.api参数
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	c.String(http.StatusOK, name+" is "+action)
	//})

	// 4.url参数
	//r.GET("/welcome", func(c *gin.Context) {
	//	name := c.DefaultQuery("name", "Jack")
	//	c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	//})

	// 5.表单参数
	//r.POST("/form", func(c *gin.Context) {
	//	// 表单参数设置默认值
	//	type1 := c.DefaultPostForm("type", "alert")
	//	// 接受其他的
	//	username := c.PostForm("username")
	//	password := c.PostForm("password")
	//	hobbys := c.PostFormArray("hobby")
	//	c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v", type1, username, password, hobbys))
	//})

	// 6.上传文件
	//r.POST("/upload", func(c *gin.Context) {
	//	file, _ := c.FormFile("file")
	//	log.Println(file.Filename)
	//	c.SaveUploadedFile(file, "./go_5/page/"+file.Filename)
	//	c.String(200, fmt.Sprintf("%s upload!", file.Filename))
	//})

	// 7.上传多文件
	//r.MaxMultipartMemory = 8 << 20
	//r.POST("/upload", func(c *gin.Context) {
	//	form, err := c.MultipartForm()
	//	if err != nil {
	//		c.String(http.StatusBadRequest, fmt.Sprintf("%s upload!", err.Error()))
	//	}
	//
	//	// 获取所有图片
	//	files := form.File["files"]
	//	for _, file := range files {
	//		// 逐一存
	//		if err := c.SaveUploadedFile(file, "./go_5/page/"+file.Filename); err != nil {
	//			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
	//		}
	//		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	//	}
	//})

	// 8.路由组
	//v1 := r.Group("v1")
	//{
	//	v1.GET("/login", login)
	//	v1.GET("submit", submit)
	//}
	//
	//v2 := r.Group("v2")
	//{
	//	v2.POST("/login", login)
	//	v2.POST("submit", submit)
	//}

	// json 绑定
	//r.POST("loginJSON", func(c *gin.Context) {
	//	// 声明接收的变量
	//	var josn Login
	//	// 将request的body中的数据，自动按照json格式解析到结构体
	//	if err := c.ShouldBindJSON(&josn); err != nil {
	//		// 返回错误信息
	//		// gin.H封装了生成json数据的工具
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if josn.User != "root" || josn.Password != "admin" {
	//		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"status": "200"})
	//})

	// 表单数据解析和绑定
	//r.POST("loginForm", func(c *gin.Context) {
	//	var form Login
	//	// Bind() 默认解析并绑定form
	//	// 根据请求头中content-type自动推断
	//	if err := c.Bind(&form); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	//		return
	//	}
	//
	//	if form.User != "root" || form.Password != "admin" {
	//		c.JSON(http.StatusBadRequest, gin.H{"status": 304})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"status": 200})
	//})

	// URI数据解析和绑定
	r.GET("/:user/:password", func(c *gin.Context) {
		var login Login
		if err := c.BindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": 200})
	})

	// 3.监听端口，默认8080
	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func getting(c *gin.Context) {

}
