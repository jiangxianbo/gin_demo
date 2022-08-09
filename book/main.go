package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("./book/templates/*")
	book := r.Group("/book")
	{
		book.GET("/list", bookListHandler)
		book.GET("/new", addBook)
		book.POST("/new", submit)
		//book.GET("/del", delook)
	}

	_ = r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	fmt.Println(bookList)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

func addBook(c *gin.Context) {
	c.HTML(http.StatusOK, "new_book.html", gin.H{})
}

func submit(c *gin.Context) {
	var book Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 添加数据库
	err := insertBook(book.Title, int(book.Price))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8000/book/list")
}

//
//func delook(c *gin.Context) {
//	c.HTML(http.StatusOK, "book_list.html", gin.H{})
//}
