package myLog

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type MyLog struct {
	level       LogLevel
	fileName    string
	filePath    string
	fileHandler *os.File
	perFileSize int64
	logChan     chan *logData
}

type logData struct {
	LineNum  int
	Msg      string
	TimeStr  string
	Level    string
	File     string
	FuncName string
}

type config struct {
	LogLevel string `cnf:"log_level"`
	FileName string `cnf:"file_name"`
	FilePath string `cnf:"file_path"`
	FileSize int    `cnf:"file_size"`
}

// 改造成channel 一步写入

var buffSize int = 5e4

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
func NewFileLog(filePath string) *MyLog {
	runtime.GOMAXPROCS(1)
	cnf := parseConf(filePath)

	l := &MyLog{
		fileName:    cnf.FileName,
		filePath:    cnf.FilePath,
		perFileSize: int64(cnf.FileSize),
		logChan:     make(chan *logData, buffSize),
	}

	l.level = l.getLogLevel(cnf.LogLevel)

	fileHandler, err := os.OpenFile(path.Join(cnf.FilePath, l.fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	l.fileHandler = fileHandler
	//go l.asyncWrite()
	return l
}

func parseConf(path string) *config {
	// 1. 打开文件
	// 2. defer file.Close()
	// 3. 读取文件内容
	// 4. 解析文件内容
	// 5. 赋值
	cnf := &config{}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	x := strings.Split(string(file), "\n")
	for _, s := range x {
		s := strings.Split(s, ":")
		setConfig(s, cnf)
	}

	return cnf
}

func setConfig(s []string, cnf *config) {
	v := reflect.ValueOf(cnf).Elem()
	t := reflect.TypeOf(*cnf)
	//fmt.Println(v.Kind(), v.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//fmt.Println(field.Name, field.Type, field.Tag.Get("cnf"))
		// k = s[0] v = s[1]
		if field.Tag.Get("cnf") == s[0] && v.Field(i).CanSet() {
			s[1] = strings.Trim(s[1], "\r ")
			// 拿到每个字段的类型
			switch field.Type.Kind() {
			case reflect.String:
				v.Field(i).SetString(s[1])
			case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
				atoi, _ := strconv.Atoi(s[1])
				v.Field(i).SetInt(int64(atoi))
			}

		}
	}
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
	now := time.Now().Format("2006-01-02 15:04:05.0000")
	funcName, file, line := RuntimeCaller()
	msg := fmt.Sprintf(format, a...)
	logData := &logData{
		LineNum:  line,
		Msg:      msg,
		TimeStr:  now,
		Level:    LevelName(level),
		File:     file,
		FuncName: funcName,
	}

	// 这样写才不会阻塞，而是丢弃
	select {
	case l.logChan <- logData:
	default:
	}
}

func (l *MyLog) asyncWrite() {
	for {
		select {
		case log := <-l.logChan:
			format := fmt.Sprintf("[%s] [%s] [%s:%s:%d]",
				log.TimeStr, log.Level, log.File, log.FuncName, log.LineNum) + log.Msg
			l.SizeCheck()
			fmt.Fprintln(l.fileHandler, format)
		}
	}
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
