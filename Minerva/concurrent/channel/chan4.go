package main

import "fmt"

// 判断channel是否关闭
// 1. value, ok := <- ch
// 2. for t := range ch{}
func main() {
	c := make(chan int, 1)
	close(c)
	v, ok := <-c
	if !ok {
		fmt.Println("channel closed")
	} else {
		fmt.Println(v)
	}

	ch := make(chan int, 100)
	go test(ch)

	for t := range ch {
		fmt.Println(t)
	}
}

func test(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	// 不关要死锁
	close(c)
}
