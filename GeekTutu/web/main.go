package main

import (
	"GeekTutu/web/gee"
	"net/http"
)

//func main() {
//	//http.HandleFunc("/", handler)
//	//log.Fatal(http.ListenAndServe(":8080", nil))
//	engine := &Engine{}
//	log.Fatal(http.ListenAndServe(":8080", engine))
//}
//
////func handler(w http.ResponseWriter, r *http.Request) {
////	w.Write([]byte("hello world"))
////}
//
//// Engine 为什么要这么写？而且ListenAndServe还要传入
//// 因为ListenAndServe需要接收一个 handler Handler
//// 而Handler 需要实现一个方法： 即 ServeHTTP(http.ResponseWriter, *http.Request)
//type Engine struct {
//}
//
//// ServeHTTP 所以这里就实现了呗。然后Engine就是实现了Handler
//func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		w.Write([]byte("base path"))
//	case "/hello":
//		w.Write([]byte("hello world"))
//	default:
//		w.WriteHeader(404)
//	}
//}

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.URL)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.URL)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
