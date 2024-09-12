package guru

import (
	"fmt"
	"reflect"
)

func reflectType(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Printf("type: %v\n", v)
}
