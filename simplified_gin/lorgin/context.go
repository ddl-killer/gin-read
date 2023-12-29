package lorgin

import (
	"fmt"
	"net/http"
	"sync"
)

type Context struct {
	w          http.ResponseWriter
	Request    *http.Request
	StatusCode int
	handlers   HandlersChain
	index      int
	fullPath   string
	mu         sync.RWMutex
	Keys       map[string]any
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w:       w,
		Request: r,
		index:   -1,
	}
}

func (c *Context) reset(w http.ResponseWriter, r *http.Request) {
	c.w = w
	c.Request = r
	c.index = -1
	c.fullPath = ""
	c.handlers = nil
	c.Keys = nil
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.w.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...any) {
	c.Status(code)
	c.w.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Get(key string) (value any, exist bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exist = c.Keys[key]
	return
}

func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Keys == nil {
		return
	}
	c.Keys[key] = value
}
