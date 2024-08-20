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
	e := gee.New()
	e.GET("/", func(e *gee.Context) {
		e.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	e.GET("/hello", func(e *gee.Context) {
		e.Json(http.StatusOK, gee.H{
			"message": "Hello Gee",
			"data":    "",
		})
	})

	e.GET("/string", func(e *gee.Context) {
		e.String(http.StatusOK, "Hello Gee >_<.%s", "   |-----|")
	})

	e.Run(":8080")
}
