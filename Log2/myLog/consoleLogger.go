package myLog

import (
	"fmt"
	"os"
	"time"
)

type ConsoleLog struct {
	level LogLevel
}

// NewConsoleLog 实例化文件项目
func NewConsoleLog(levelStr string) *ConsoleLog {
	l := &ConsoleLog{}
	l.level = l.getLogLevel(levelStr)
	return l
}

func (l *ConsoleLog) write(level LogLevel, format string, a ...interface{}) {
	if l.level >= level {
		return
	}

	// 要包含时间，日志级别，调用的文件，调用的函数，信息
	now := time.Now().Format("2006-01-02 15:04:05.0000")

	funcName, file, line := RuntimeCaller()

	format = fmt.Sprintf("[%s] [%s] [%s:%s:%d]", now, LevelName(level), file, funcName, line) + format
	logStr := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stdout, logStr)
}

func (l *ConsoleLog) Debug(format string, a ...interface{}) {
	l.write(DEBUG, format, a...)
}

func (l *ConsoleLog) Trace(format string, a ...interface{}) {
	l.write(TRACE, format, a...)
}

func (l *ConsoleLog) Info(format string, a ...interface{}) {
	l.write(INFO, format, a...)
}

func (l *ConsoleLog) Warn(format string, a ...interface{}) {
	l.write(WARNING, format, a...)
}

func (l *ConsoleLog) Error(format string, a ...interface{}) {
	l.write(ERROR, format, a...)
}

func (l *ConsoleLog) Fatal(format string, a ...interface{}) {
	l.write(FATAL, format, a...)
}

func (l *ConsoleLog) Close() {
}
