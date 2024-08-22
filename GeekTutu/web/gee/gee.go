package gee

import (
	"net/http"
)

// HandlerFunc 重新定义了context不是http.HandlerFunc了
type HandlerFunc func(*Context)

// Engine 定义一个引擎
type Engine struct {
	router *router
}

// 实现ServeHTTP方法
// 路由接管了解析和操作
// 这里是每个请求都会到这里来
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cxt := NewContext(w, r)
	e.router.handle(cxt)
}

// New 实例化Engine
func New() *Engine {
	return &Engine{router: newRouter()} // newRouter 也可以写成 &router{make(map[string]HandlerFunc)}
}

// 路由添加的方法
func (e *Engine) addRoute(uri string, handler HandlerFunc, method string) {
	e.router.AddRoute(uri, handler, method)
}

// POST 请求方式的添加方法
func (e *Engine) POST(uri string, handler HandlerFunc) {
	e.router.AddRoute(uri, handler, "POST")
}

// GET add get router
func (e *Engine) GET(uri string, handler HandlerFunc) {
	e.router.AddRoute(uri, handler, "GET")
}

// Run 执行
func (e *Engine) Run(serveAddr string) error {
	return http.ListenAndServe(serveAddr, e)
}
