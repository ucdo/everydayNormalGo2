package myLog

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

var logger MyLogger

func TestSizeCheck(t *testing.T) {
	t.Helper()
	logx := NewFileLog("DEBUG", "../runtime", "2024-09-11_1.log")
	//go func() {
	//	wg.Add(1)
	//	defer wg.Done()
	//	logx.SizeCheck()
	//}()
	//wg.Wait()
	logx.SizeCheck()
	//for i := 0; i < 1e5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		logx.Debug("用户%10d正在疯狂尝试登录", i)
	//	}(i)
	//}
	//wg.Wait()
}

func TestConsoleOutput(t *testing.T) {
	t.Helper()
	logger = NewFileLog("DEBUG", "../runtime", "2024-09-11_1.log")
	logger.Debug("用户%10d正在疯狂尝试登录", 123123)
	logger.Info("用户%10d正在疯狂尝试登录", 123123)
	logger.Trace("用户%10d正在疯狂尝试登录", 123123)
	logger.Warn("用户%10d正在疯狂尝试登录", 123123)
	logger.Error("用户%10d正在疯狂尝试登录", 123123)
	logger.Fatal("用户%10d正在疯狂尝试登录", 123123)

	logger = NewConsoleLog("DEBUG")

	logger.Debug("用户%10d正在疯狂xx", 123123)
	logger.Info("用户%10d正在疯狂xx", 123123)
	logger.Trace("用户%10d正在疯狂xx", 123123)
	logger.Warn("用户%10d正在疯狂xx", 123123)
	logger.Error("用户%10d正在疯狂xx", 123123)
	logger.Fatal("用户%10d正在疯狂xx", 123123)
}
