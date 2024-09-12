package main

import (
	"Log2/myLog"
	"sync"
	"time"
)

var wg sync.WaitGroup

var Logger myLog.MyLogger

func main() {
	// call
	now := time.Now().Format("2006-01-02") + "_1.log"
	logger := myLog.NewFileLog("debug", "./runtime/", now)

	for i := 0; i < 1e6; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger.Warn("写写warn")
			logger.Fatal("用户%4d正在疯狂尝试登录", i)
		}(i)
	}
	wg.Wait()
}
