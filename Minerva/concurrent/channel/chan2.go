package main

import "fmt"

// 无缓冲通道
// 无缓冲通道相当于是同步操作
// 同一个通道，必须有人收，有人发，不然就一直阻塞 = deadlock
func noBuffer(c chan int) {
	x := <-c
	// 这里有一个比较有意思的是，下面的代码有可能不会输出
	// 因为 <-c 取完值之后channel就非阻塞了
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go noBuffer(c)
	c <- 1
	fmt.Println("main 结束了")
}
