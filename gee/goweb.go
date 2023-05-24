package gee

import (
	"net/http"
)

type HandleFunc func(ctx *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method, pattern string, handle HandleFunc) {
	e.router.addRouter(method, pattern, handle)
}

func (e *Engine) GET(pattern string, handle HandleFunc) {
	e.addRoute("GET", pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandleFunc) {
	e.addRoute("POST", pattern, handle)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
