package guru

import (
	"fmt"
	"reflect"
)

type rf struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Printf("type: %v\n", t)
	fmt.Printf("value: %v\n", v)
}

// 不能这样改
func reflectModify(a interface{}) {
	a = 100
}

func reflectModify2(a interface{}) {
	v := reflect.ValueOf(a)
	kind := v.Kind()
	switch kind {
	case reflect.Ptr:
		v.Elem().SetInt(100)
	}
}

func rfStruct() {

}
