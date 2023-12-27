package lorgin

import (
	"fmt"
	"net/http"
)

type Context struct {
	w          http.ResponseWriter
	r          *http.Request
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w: w,
		r: r,
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
