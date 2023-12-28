package main

import (
	"github.com/ddl-killer/gin-read/simplified_gin/lorgin"
)

func main() {
	e := lorgin.Default()
	e.Handle("GET", "/hello", func(c *lorgin.Context) {
		c.String(200, "hello, %s", "lor")
	})
	emailGroup := e.Group("email")
	{
		emailGroup.GET("/narvar", func(c *lorgin.Context) {
			c.String(403, "403")
		})
		emailGroup.GET("/shopify", func(c *lorgin.Context) {
			c.String(200, "200")
		})
	}
	e.GET("/panic", func(c *lorgin.Context) {
		panic("panic")
	})
	if err := e.Run(); err != nil {
		panic(err.Error())
	}
}
