package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	//r.DELETE("/delete", func(c *gin.Context) {
	//	c.String(200, "delete")
	//})
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	r.Any("/any", func(c *gin.Context) { // Any function matches all 8 types of requests
		c.String(200, "any")
	})
	r.Run()
}
