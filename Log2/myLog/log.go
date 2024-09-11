package myLog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

const (
	DEBUG   = iota //测试的日志
	TRACE          //链路日志
	INFO           //信息日志
	WARNING        //警告日志
	ERROR          //错误日志
	FATAL          // 严重错误
)

var mtx sync.Mutex

func LevelName(level int) string {
	mp := map[int]string{
		DEBUG:   "DEBUG",
		TRACE:   "TRACE",
		INFO:    "INFO",
		WARNING: "WARNING",
		ERROR:   "ERROR",
		FATAL:   "FATAL",
	}

	if _, ok := mp[level]; !ok {
		level = DEBUG
	}

	return mp[level]
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

type MyLog struct {
	level       int
	fileName    string
	filePath    string
	fileHandler *os.File
	perFileSize int64
}

// NewMyLog 实例化文件项目
func NewMyLog(level int, filePath, fileName string) *MyLog {
	l := &MyLog{
		level:       level,
		fileName:    fileName,
		filePath:    filePath,
		perFileSize: 1024 * 1024 * 20,
	}

	fileHandler, err := os.OpenFile(path.Join(filePath, l.fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	l.fileHandler = fileHandler
	return l
}

func (l *MyLog) genLogFileName() string {
	fileInfo, err := l.fileHandler.Stat()
	if err != nil {
		panic(err)
	}

	if fileInfo.Size() < l.perFileSize {
		return l.fileName
	}

	folder, err := os.ReadDir(l.filePath)
	if err != nil {
		panic(err)
	}
	count := 0
	for _, entry := range folder {
		if !entry.IsDir() {
			count++
		}
	}
	now := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s_%d.log", now, count+1)
}

func (l *MyLog) SizeCheck() {
	mtx.Lock()
	defer mtx.Unlock()

	fileInfo, err := l.fileHandler.Stat()
	if err != nil {
		panic(err)
	}

	if fileInfo.Size() >= l.perFileSize {
		// 创建一个新的文件
		l.fileName = l.genLogFileName()
		fileHandler, err := os.OpenFile(path.Join(l.filePath, l.fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		// 关闭久的文件
		l.fileHandler.Close()
		l.fileHandler = fileHandler
	}
}

func (l *MyLog) genDefault(format string) string {
	// 要包含时间，日志级别，调用的文件，调用的函数，信息
	now := time.Now().Format("2006-01-02 15:04:05.0000")
	level := LevelName(l.level)
	funcName, file, line := RuntimeCaller()

	format = fmt.Sprintf("[%s] [%s] [%s:%s:%d]", now, level, file, funcName, line) + format + "\n"
	return format
}

func (l *MyLog) Debug(format string, a ...any) {
	if l.level < WARNING {
		return
	}

	l.SizeCheck()
	format = l.genDefault(format)
	fmt.Fprintf(l.fileHandler, format, a...)
}
