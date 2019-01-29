package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main()  {
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	
	r.GET("/echo/:words", func(c *gin.Context) {
		words := c.Param("words")
		c.String(200, "%s", words)
	})
	
	r.GET("/echo/:words/*action", func(c *gin.Context) {
		words := c.Param("words")
		action := c.Param("action")
		
		msg := words + " will be " + action
		
		c.String(http.StatusOK, msg)
	})
	
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("fname", "Guest")
		lastname := c.Query("lname")
		
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	
	r.POST("/form", func(c *gin.Context) {
		msg := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		
		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"message": msg,
			"nick": nick,
		})
	})
	
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		msg := c.PostForm("message")
		
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, msg)
	})
	
	r.POST("/post/map", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		
		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	
	
	r.Run()
	
}
