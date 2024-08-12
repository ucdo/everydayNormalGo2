package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) hello(name string) {
	fmt.Printf("hello %s,I am %s\n", name, s.name)
}

func main() {
	// 即可以通过实例名，student调用student上
	s := &student{name: "Tom"}
	s.hello("jack")
}
