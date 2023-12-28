package lorgin

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		c.Next()

		method := c.Request.Method
		path := c.Request.URL
		statusCode := c.StatusCode
		log.Println("method:", method, "path:", path, "statusCode:", statusCode, "cost:", time.Since(start))
	}
}
