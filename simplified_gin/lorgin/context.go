package lorgin

import (
	"fmt"
	"net/http"
)

type Context struct {
	w          http.ResponseWriter
	Request    *http.Request
	StatusCode int
	handlers   HandlersChain
	index      int
	fullPath   string
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w:       w,
		Request: r,
		index:   -1,
	}
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.w.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...any) {
	c.Status(code)
	c.w.Write([]byte(fmt.Sprintf(format, values)))
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}
