package main

import (
	"Log2/myLog"
	"sync"
)

var wg sync.WaitGroup

var Logger myLog.MyLogger

func main() {
	logger := myLog.NewFileLog("./main.cnf")

	for i := 0; i < 1e5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger.Warn("写写warn")
			logger.Fatal("用户%4d正在疯狂尝试登录", i)
		}(i)
	}
	wg.Wait()
}
