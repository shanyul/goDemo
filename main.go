package main

import (
	router "gee/routers"
	"net/http"
)

func main() {
	r := router.New()
	r.GET("/", func(c *router.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *router.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *router.Context) {
		c.JSON(http.StatusOK, router.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8080")
}
