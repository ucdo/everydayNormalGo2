package main

import "fmt"

// select 多路复用
// 多个通道可以同时操作
/*
for{
	select{
	case <- ch:
		...
	case <- ch2:
		....
	default:
		...
	}
}
*/

func f1(ch chan string) {
	for i := '0'; i < '9'; i++ {
		ch <- string(i)
	}
}

func f2(ch chan string) {
	for i := 'a'; i < 'j'; i++ {
		ch <- string(i)
	}
}

func main() {
	ch := make(chan string, 1)
	go f1(ch)
	go f2(ch)
	for {
		select {
		case v := <-ch:
			fmt.Printf("%s\n", v)
		}
	}
}
