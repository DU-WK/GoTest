package KGin

import (
	"net/http"
)

type Engine struct {
	router *Router
}

func NewEngine() *Engine {
	return &Engine{
		router: &Router{
			routes: make(map[string]HandlerFunc),
		},
	}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	engine.router.HandleFunc(context)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) AddRoute(method string, url string, handler HandlerFunc) {
	routerKey := method + "-" + url
	engine.router.routes[routerKey] = handler
}

func (engine *Engine) GET(url string, handler HandlerFunc) {
	routerKey := "GET" + "-" + url
	engine.router.routes[routerKey] = handler
}

func (engine *Engine) POST(url string, handler HandlerFunc) {
	routerKey := "POST" + "-" + url
	engine.router.routes[routerKey] = handler
}

func (engine *Engine) PUT(url string, handler HandlerFunc) {
	routerKey := "PUT" + "-" + url
	engine.router.routes[routerKey] = handler
}

func (engine *Engine) DELETE(url string, handler HandlerFunc) {
	routerKey := "DELETE" + "-" + url
	engine.router.routes[routerKey] = handler
}
