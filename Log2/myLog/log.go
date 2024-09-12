package myLog

import (
	"path"
	"runtime"
	"strings"
	"sync"
)

type MyLogger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Close()
}

type LogLevel int

const (
	DEBUG   LogLevel = iota //测试的日志
	TRACE                   //链路日志
	INFO                    //信息日志
	WARNING                 //警告日志
	ERROR                   //错误日志
	FATAL                   // 严重错误
)

var mtx sync.Mutex

func LevelName(level LogLevel) string {
	mp := map[LogLevel]string{
		DEBUG:   "DEBUG",
		TRACE:   "TRACE",
		INFO:    "INFO",
		WARNING: "WARN",
		ERROR:   "ERROR",
		FATAL:   "FATAL",
	}

	if _, ok := mp[level]; !ok {
		level = DEBUG
	}

	return mp[level]
}

func (l *ConsoleLog) getLogLevel(levelStr string) LogLevel {
	level := strings.ToUpper(levelStr)
	switch level {
	case "DEBUG":
		return DEBUG
	case "TRACE":
		return TRACE
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return DEBUG
	}
}

func RuntimeCaller() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		panic("runtime caller error")
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
