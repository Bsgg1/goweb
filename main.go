package main

import (
	"gee"
	"net/http"
)

func main() {
	e := gee.New()
	e.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>杨恩文大傻逼</h1>")
	})
	e.GET("/hello/:name", func(c *gee.Context) {

		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	e.Run(":9999")

}
