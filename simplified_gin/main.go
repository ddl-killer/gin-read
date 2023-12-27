package main

import (
	"fmt"
	"net/http"

	"github.com/ddl-killer/gin-read/simplified_gin/lorgin"
)

func main() {
	e := lorgin.Default()
	e.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello lor")
	})
	if err := e.Run(); err != nil {
		panic(err.Error())
	}
}
