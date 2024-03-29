package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user"  xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password"  xml:"password" binding:"required"`
}

// MiddleWare 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
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
	//r.GET("/:user/:password", func(c *gin.Context) {
	//	var login Login
	//	if err := c.BindUri(&login); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	//		return
	//	}
	//	if login.User != "root" || login.Password != "admin" {
	//		c.JSON(http.StatusBadRequest, gin.H{"status": 304})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{"status": 200})
	//})

	// 各种数据格式的响应
	//// 1.json
	//r.GET("/someJSON", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	//})
	//
	//// 2.结构体响应
	//r.GET("/someStruct", func(c *gin.Context) {
	//	var msg struct {
	//		Name    string
	//		Message string
	//		Number  int
	//	}
	//	msg.Name = "root"
	//	msg.Message = "message"
	//	msg.Number = 123
	//	c.JSON(200, msg)
	//})
	//
	//// 3.XML响应
	//r.GET("/someXML", func(c *gin.Context) {
	//	c.XML(200, gin.H{"message": "abc"})
	//})
	//
	//// 4.YAML响应
	//r.GET("/someYAML", func(c *gin.Context) {
	//	c.YAML(200, gin.H{"name": "zhangsan"})
	//})
	//// 5.protobuf格式,谷歌开发的高效存储读取的工具
	//// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	//r.GET("/someProtobuf", func(c *gin.Context) {
	//	reps := []int64{int64(1), int64(1)}
	//	// 定义数据
	//	label := "label"
	//	// 传protobuf格式
	//	data := &protoexample.Test{
	//		Label: &label,
	//		Reps:  reps,
	//	}
	//	c.ProtoBuf(200, data)
	//})

	// HTML模板渲染
	// 加载文件
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/index", func(c *gin.Context) {
	//	// 根据文件名
	//	c.HTML(200, "index.tmpl", gin.H{"title": "我的标题"})
	//})

	// 重定向
	//r.GET("/redirect", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	//})

	//// 1.异步
	//r.GET("/long_async", func(c *gin.Context) {
	//	// 需要搞一个副本
	//	copyContext := c.Copy()
	//	// 异步处理
	//	go func() {
	//		time.Sleep(time.Second * 3)
	//		log.Println("异步执行", copyContext.Request.URL.Path)
	//	}()
	//})
	//
	//// 2.同步
	//r.GET("/long_sync", func(c *gin.Context) {
	//	time.Sleep(time.Second * 3)
	//	log.Println("异步执行", c.Request.URL.Path)
	//})

	// 注册中间件
	//r.Use(MiddleWare())
	//{
	//	r.GET("middleware", func(c *gin.Context) {
	//		// 取值
	//		req, _ := c.Get("request")
	//		fmt.Println("request:", req)
	//		// 页面返回
	//		c.JSON(200, gin.H{"request": req})
	//	})
	//
	//	r.GET("middleware2", MiddleWare(), func(c *gin.Context) {
	//		// 取值
	//		req, _ := c.Get("request")
	//		fmt.Println("request:", req)
	//		// 页面返回
	//		c.JSON(200, gin.H{"request": req})
	//	})
	//}

	// 中间件联系
	//r.Use(myTime)
	//group := r.Group("/shopping")
	//{
	//	group.GET("/index", shopIndexHandler)
	//	group.GET("/home", shopHomeHandler)
	//}

	// 服务端要给客户端cookie
	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置cookie
			// maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			// secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "127.0.0.1", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})

	// 3.监听端口，默认8080
	r.Run(":8000")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(time.Second * 5)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(time.Second * 3)
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
