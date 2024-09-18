package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello world")
}

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}

func main() {
	// 设置go程序只用一个逻辑核心
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
