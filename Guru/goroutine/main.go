package main

import "fmt"

func main() {
	var a = 1
	fmt.Println(a)
	ch := make(chan interface{})
	go func() {
		ch <- "xxxx"
	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	}
}
