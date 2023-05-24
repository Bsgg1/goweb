package gee

import "net/http"

type Router struct {
	handlers map[string]HandleFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandleFunc)}
}

func (r *Router) addRouter(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handle
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}
