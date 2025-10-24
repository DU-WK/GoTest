package KGin

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]http.HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]http.HandlerFunc),
	}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	strMethod := r.Method
	strURLPath := r.URL.Path
	routerKey := strMethod + "-" + strURLPath

	fmt.Println(routerKey)

	if handler, ok := engine.router[routerKey]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (engine *Engine) AddRoute(method string, url string, handler http.HandlerFunc) {
	routerKey := method + "-" + url
	engine.router[routerKey] = handler
}

func (engine *Engine) GET(url string, handler http.HandlerFunc) {
	routerKey := "GET" + "-" + url
	engine.router[routerKey] = handler
}

func (engine *Engine) POST(url string, handler http.HandlerFunc) {
	routerKey := "POST" + "-" + url
	engine.router[routerKey] = handler
}

func (engine *Engine) PUT(url string, handler http.HandlerFunc) {
	routerKey := "PUT" + "-" + url
	engine.router[routerKey] = handler
}

func (engine *Engine) DELETE(url string, handler http.HandlerFunc) {
	routerKey := "DELETE" + "-" + url
	engine.router[routerKey] = handler
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}
