package main

import "fmt"

func main() {
	// 1. 定义一个chan 传递int
	var ch1 chan int
	// 2. 定义一个chan 传递string
	var ch2 chan string
	// 打印
	fmt.Println(ch1, ch2)
	// notice: 往上面的channel <- 会报错。因为没定义缓冲区，没地方写

	// channel 是引用类型，用make，零值是nil
	ch := make(chan int, 1)
	// 写入
	ch <- 1
	// 读取并保存
	x := <-ch
	fmt.Println("x:", x)
	// 写入超过buffer容量的会被阻塞,并发出fatal error,程序中断
	ch <- 1
	ch <- 2
	// 关闭channel
	close(ch)

	// 关闭之后再读取，不会panic
	//如果关闭之前有值，则读取相应的值，否则返回对应的零值
	fmt.Println(<-ch)

	// 往关闭了的channel写入，panic
	ch <- 1
}
