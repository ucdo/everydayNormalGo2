package main

import "fmt"

func main() {
	ch := make(chan interface{})
	go func() {
		ch <- "xxxx"
	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	}
}
