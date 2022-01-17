package main

import (
	"gkz"
	"net/http"
)

func main() {
	r := gkz.New()
	r.GET("/", func(c *gkz.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gkz</h1>")
	})
	r.GET("/hello", func(c *gkz.Context) {
		// except /hello?name=zhy
		c.String(http.StatusOK, "hello %s,you'r at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gkz.Context) {
		c.JSON(http.StatusOK, gkz.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
