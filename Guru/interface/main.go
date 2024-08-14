package main

import (
	Qeueu "Guru/interface/queue"
	"fmt"
	"io"
	"net/http"
)

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

func getPerson() person {
	return &student{}
}

func main() {
	//var s = getPerson()
	//
	//fmt.Println(s.getName())
	//var s = getImpl()
	//fmt.Println(s.get("https://baidu.com"))
	q := &Qeueu.Queue{}
	fmt.Println(q.IsEmpty())
	q.Push(1)
	q.Push(2)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}

type implementer interface {
	get(string) string
}

type fake struct {
}

func (f fake) get(s string) string {
	return "fake get."
}

type real struct {
}

func (r real) get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func getImpl() implementer {
	return &real{}
}
