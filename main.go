package main

import (
	"fmt"
	router "gee/routers"
	"log"
	"net/http"
	"text/template"
	"time"
)

func onlyForV2() router.HandlerFunc {
	return func(c *router.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))

	}
}

type student struct {
	Name string
	Age  string
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := router.New()
	r.Use(router.Logger())

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/assert", "./static")

	stu1 := &student{Name: "hademen", Age: "18"}
	stu2 := &student{Name: "wandefa", Age: "30"}

	r.GET("/", func(c *router.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *router.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", router.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *router.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", router.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *router.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *router.Context) {
			c.JSON(http.StatusOK, router.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.Run(":8080")
}
