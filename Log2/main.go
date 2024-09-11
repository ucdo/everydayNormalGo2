package main

import (
	"Log2/myLog"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// call
	now := time.Now().Format("2006-01-02") + "_1.log"
	logx := myLog.NewMyLog(myLog.DEBUG, "./runtime/", now)

	for i := 0; i < 1e5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logx.Debug("用户%10d正在疯狂尝试登录", i)
		}(i)
	}
	wg.Wait()
}
