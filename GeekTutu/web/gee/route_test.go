package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	//r.AddRoute("/", nil, "GET")
	r.AddRoute("/hello/:name", nil, "GET")
	//r.AddRoute("/hello/b/c", nil, "GET")
	//r.AddRoute("/hi/:name", nil, "GET")
	//r.AddRoute("/assets/*filepath", nil, "GET")
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	//ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"}) // false
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name", "*"}) // false
	if !ok {
		t.Fatal("test parse pattern /p/*name/* failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, params := r.getRoute("/hello/geektutu", "GET")
	if n == nil {
		t.Fatal("test get route failed")
	}

	if n.pattern != "hello/:name" {
		t.Fatal("test get route failed")
	}

	if params["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, params["name"])
}
