package main

import (
	"gee"
)

func main() {
	e := gee.New()
	e.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>杨恩文大傻逼</h1>")
	})

	e.Run(":9999")

}
