package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("login_cookie"); err == nil {
			if cookie == "value_cookie" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("login_cookie", "value_cookie", 60, "/", "127.0.0.1", false, true)
		c.String(http.StatusOK, "Login success")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8000")
}
