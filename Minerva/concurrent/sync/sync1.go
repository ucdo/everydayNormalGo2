package main

import (
	"fmt"
	"sync"
)

// 资源竞争

var x int64
var wg sync.WaitGroup
var mtx sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		mtx.Lock()
		// 加锁因为会有竞争
		// 多个goroutine等待时，随机唤醒一个
		x += 1
		mtx.Unlock()
	}
}

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		add()
	}()
	go func() {
		defer wg.Done()
		add()
	}()
	wg.Wait()
	fmt.Println(x)
	return
}
