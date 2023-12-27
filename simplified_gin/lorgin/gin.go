package lorgin

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router map[string]HandlerFunc
}

func Default() *Engine {
	return New()
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	key := generateRouterKey(method, path)
	c := newContext(w, r)
	if h := e.router[key]; h != nil {
		h(c)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func (e *Engine) Handle(method string, relativePath string, handler HandlerFunc) {
	key := generateRouterKey(method, relativePath)
	if _, ok := e.router[key]; ok {
		panic("already have")
	} else {
		e.router[key] = handler
	}
}

func generateRouterKey(method string, relativePath string) string {
	return fmt.Sprintf("%s-%s", method, relativePath)
}

func (e *Engine) GET(relativePath string, handler HandlerFunc) {
	e.Handle("GET", relativePath, handler)
}

func (e *Engine) POST(relativePath string, handler HandlerFunc) {
	e.Handle("POST", relativePath, handler)
}

func (e *Engine) Run(addr ...string) error {
	address := resolveAddress(addr)
	if err := http.ListenAndServe(address, e); err != nil {
		return err
	}
	return fmt.Errorf("run failed")
}
