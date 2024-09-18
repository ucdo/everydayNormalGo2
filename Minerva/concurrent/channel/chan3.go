package main

import "fmt"

// 带缓冲的通道

func buffer(c chan int) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan int, 3)
	c <- 1
	fmt.Println(len(c), cap(c))
	go buffer(c)
	c <- 2
	fmt.Println(len(c), cap(c))
}
