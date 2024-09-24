package main

import (
	"fmt"
	"sync"
)

var one sync.Once // 只执行一次

var mp map[string]int

func loadMap() {
	mp = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
}

func main() {
	if mp == nil {
		one.Do(loadMap) // 会维护一个缓存，告诉其他调用这个方法的人，这个方法有没有执行
	}
	fmt.Println(mp["a"])
}
