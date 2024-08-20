package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 暂时没看明白是在做什么用的
type H map[string]interface{}

type Context struct {
	W      http.ResponseWriter
	R      *http.Request
	URL    string
	Method string
	Status int
}

// NewContext 新建上下文。按理来说是每个请求新建一个上下文
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		URL:    r.URL.Path,
		Method: r.Method,
	}
}

// PostFrom 从post form 获取数据
func (c *Context) PostFrom(key string) string {
	return c.R.FormValue(key)
}

// Query 从query里面获取数据、
func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

// SetStatus 设置响应状态
func (c *Context) SetStatus(code int) {
	c.Status = code
	c.W.WriteHeader(code)
}

// SetHeader 设置值头
func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)
}

// String 返回string类型的数据
func (c *Context) String(code int, form string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.W.Write([]byte(fmt.Sprintf(form, value...)))
}

// Json 返回json格式的数据
func (c *Context) Json(code int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)

	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(data); err != nil {
		c.SetStatus(500)
		c.W.Write([]byte(err.Error()))
	}
}

// Data 不晓得啥子类型的数据.e好像是二进制
func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.W.Write(data)
}

// HTML 返回HTML格式的数据
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.W.Write([]byte(html))
}
