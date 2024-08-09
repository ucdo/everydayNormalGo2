package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "xxx"
	s2 := "xx中文"
	fmt.Println("type is: ", reflect.TypeOf(s1).Kind())
	fmt.Println("type is: ", reflect.TypeOf(s2))
}
