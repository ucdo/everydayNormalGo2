package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (r *router) AddRoute(url string, handler HandlerFunc, method string) {
	log.Printf("Route %4s - %s", method, url)
	key := method + "_" + url
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "_" + c.URL
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 page not found: %s", c.URL)
	}
}
