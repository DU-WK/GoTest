package KGin

import "net/http"

type HandlerFunc func(context *Context)

type Router struct {
	routes map[string]HandlerFunc
}

func (router *Router) HandleFunc(c *Context) {
	routerKey := c.Method + "-" + c.Path
	if handler, ok := router.routes[routerKey]; ok {
		handler(c)
	} else {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte("404 Not Found"))
	}
}
