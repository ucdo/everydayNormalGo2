package main

import "fmt"

type person interface {
	getName() string
}

type student struct {
	name string
	age  int
}

type worker struct {
	name string
	age  int
}

func (s *student) getName() string {
	return s.name
}

func (w *worker) getName() string {
	return w.name
}

func main() {
	var s person = &student{
		name: "xx",
	}

	fmt.Println(s.getName())
}
