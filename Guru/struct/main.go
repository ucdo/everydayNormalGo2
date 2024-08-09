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
	s := &student{name: "Tom"}
	s.hello("jack")
}
