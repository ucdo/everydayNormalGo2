package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁： 读比写多很多，则可以用这个提高性能
var rwLock sync.RWMutex

func read() {
	fmt.Println(x)
	time.Sleep(time.Millisecond * 10)
}

func write() {
	x += 1
	time.Sleep(time.Millisecond * 50)
}

func main() {
	
}
