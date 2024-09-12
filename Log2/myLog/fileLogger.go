package myLog

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type MyLog struct {
	level       LogLevel
	fileName    string
	filePath    string
	fileHandler *os.File
	perFileSize int64
}

func (l *MyLog) getLogLevel(levelStr string) LogLevel {
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

// NewFileLog 实例化文件项目
func NewFileLog(levelStr string, filePath, fileName string) *MyLog {
	l := &MyLog{
		fileName:    fileName,
		filePath:    filePath,
		perFileSize: 1024 * 1024 * 20,
	}

	l.level = l.getLogLevel(levelStr)

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
		l.CreateFile()
	}
}

func (l *MyLog) CreateFile() {
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

func (l *MyLog) write(level LogLevel, format string, a ...interface{}) {
	if l.level >= level {
		return
	}

	l.SizeCheck()
	// 要包含时间，日志级别，调用的文件，调用的函数，信息
	now := time.Now().Format("2006-01-02 15:04:05.0000")

	funcName, file, line := RuntimeCaller()

	format = fmt.Sprintf("[%s] [%s] [%s:%s:%d]", now, LevelName(level), file, funcName, line) + format
	logStr := fmt.Sprintf(format, a...)
	fmt.Fprintln(l.fileHandler, logStr)
}

func (l *MyLog) Debug(format string, a ...interface{}) {
	l.write(DEBUG, format, a...)
}

func (l *MyLog) Trace(format string, a ...interface{}) {
	l.write(TRACE, format, a...)
}

func (l *MyLog) Info(format string, a ...interface{}) {
	l.write(INFO, format, a...)
}

func (l *MyLog) Warn(format string, a ...interface{}) {
	l.write(WARNING, format, a...)
}

func (l *MyLog) Error(format string, a ...interface{}) {
	l.write(ERROR, format, a...)
}

func (l *MyLog) Fatal(format string, a ...interface{}) {
	l.write(FATAL, format, a...)
}

func (l *MyLog) Close() {
	l.fileHandler.Close()
}
