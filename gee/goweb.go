package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (e *Engine) addRouter(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handle
}

func (e *Engine) GET(pattern string, handle HandleFunc) {
	e.addRouter("GET", pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandleFunc) {
	e.addRouter("POST", pattern, handle)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
