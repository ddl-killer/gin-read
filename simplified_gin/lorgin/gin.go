package lorgin

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type HandlersChain []HandlerFunc

type Engine struct {
	RouterGroup
	trees methodTrees
}

func Default() *Engine {
	return New()
}

func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		trees: make(methodTrees, 0, 9),
	}
	engine.RouterGroup.engine = engine
	return engine
}

func (e *Engine) addRoute(httpMethod, path string, handlers HandlersChain) {
	root := e.trees.get(httpMethod)
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		e.trees = append(e.trees, methodTree{
			method: httpMethod,
			root:   root,
		})
	}
	root.addRoute(path, handlers)
}

func (e *Engine) Use(middleware ...HandlerFunc) IRoutes {
	e.RouterGroup.Use(middleware...)
	return e
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.handleHTTPRequest(c)
}

func (e *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	t := e.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		value := root.getValue(rPath)
		if value.handlers != nil {
			c.handlers = value.handlers
			c.fullPath = value.fullPath
			c.Next()
			return
		}
	}
	c.String(404, "404 NOT FOUND: %s\n", c.Request.URL)
}

func (e *Engine) Run(addr ...string) error {
	address := resolveAddress(addr)
	if err := http.ListenAndServe(address, e); err != nil {
		return err
	}
	return fmt.Errorf("run failed")
}
