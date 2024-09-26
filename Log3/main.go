package main

import (
	"Log2/myLog"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var Logger myLog.MyLogger

func main() {
	now := time.Now()
	logger := myLog.NewFileLog("./main.cnf")
	defer logger.Close()

	for i := 0; i < 1e5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger.Warn("写写warn")
			logger.Fatal("用户%4d正在疯狂尝试登录", i)
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(now))
}
