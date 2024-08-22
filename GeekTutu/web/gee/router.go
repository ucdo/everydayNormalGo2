package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		make(map[string]*node),
		make(map[string]HandlerFunc),
	}
}

func ParsePattern(pattern string) []string {
	result := make([]string, 0)
	parts := strings.Split(pattern, "/")
	for _, part := range parts {
		if part == "" {
			continue
		}

		result = append(result, part)
		if part == "*" {
			break
		}
	}
	// /a/*/b => [a] ?
	return result
}

func parsePattern(pattern string) []string {
	result := make([]string, 0)
	parts := strings.Split(pattern, "/")
	for _, part := range parts {
		if part == "" {
			continue
		}

		result = append(result, part)
		if part == "*" {
			break
		}
	}
	// /a/*/b => [a] ?
	return result
}

// AddRoute /a/*/b 的handler实际上是挂在node a这节点上的 /*/b 都没匹配了
func (r *router) AddRoute(url string, handler HandlerFunc, method string) {
	log.Printf("Route %4s - %s", method, url)
	parts := parsePattern(url)
	key := method + "_" + url
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(url, parts, 0)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.URL, c.Method)
	if n != nil {
		c.Params = params
		key := c.Method + "_" + c.URL
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 page not found: %s", c.URL)
	}
}

func (r *router) getRoute(url string, method string) (n *node, params map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return
	}

	parts := parsePattern(url)
	params = make(map[string]string)

	n = root.search(parts, 0)
	if n == nil {
		return
	}

	tempParts := parsePattern(n.pattern)
	for i, part := range tempParts {
		if part[0] == ':' {
			params[part[1:]] = parts[i]
		}

		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(parts[i:], "/")
			break
		}
	}

	return

}
