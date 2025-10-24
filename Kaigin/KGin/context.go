package KGin

import "net/http"

type Context struct {
	Method  string
	Path    string
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Method:  r.Method,
		Path:    r.URL.Path,
		Writer:  w,
		Request: r,
	}
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.Request.PostFormValue(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}
