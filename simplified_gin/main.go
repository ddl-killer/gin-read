package main

import (
	"github.com/ddl-killer/gin-read/simplified_gin/lorgin"
)

func main() {
	e := lorgin.Default()
	e.Handle("GET", "/hello", func(c *lorgin.Context) {
		c.String(200, "hello, %s", "lor")
	})
	if err := e.Run(); err != nil {
		panic(err.Error())
	}
}
