package myLog

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MyLog struct {
	level       LogLevel
	fileName    string
	filePath    string
	fileHandler *os.File
	perFileSize int64
}

type config struct {
	LogLevel string `cnf:"log_level"`
	FileName string `cnf:"file_name"`
	FilePath string `cnf:"file_path"`
	FileSize int    `cnf:"file_size"`
}

// 改造成channel 一步写入

var buffSize int = 5e4
var ch = make(chan string, buffSize) // 定义一个缓冲区为5万的channel
var mp = sync.Map{}

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
	}

	l.level = l.getLogLevel(cnf.LogLevel)

	fileHandler, err := os.OpenFile(path.Join(cnf.FilePath, l.fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	l.fileHandler = fileHandler
	go l.syncWrite()
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

	// 在这里等，5万条了再写入
	// 要包含时间，日志级别，调用的文件，调用的函数，信息
	now := time.Now().Format("2006-01-02 15:04:05.0000")

	funcName, file, line := RuntimeCaller()

	format = fmt.Sprintf("[%s] [%s] [%s:%s:%d]", now, LevelName(level), file, funcName, line) + format
	logStr := fmt.Sprintf(format, a...)
	key := fmt.Sprintf("%d_%d", time.Now().UnixNano(), rand.Int31n(1e6))
	mp.Store(key, logStr)
	ch <- key
	//fmt.Fprintln(l.fileHandler, logStr)
}

func (l *MyLog) syncWrite() {
	var logs []string
	var writeCount = buffSize
	for {
		select {
		case key := <-ch:
			logs = append(logs, key)

			if len(logs) >= writeCount {
				mtx.Lock()
				tmp := logs[:writeCount]
				logs = logs[writeCount:]
				mtx.Unlock()

				// 写入。清空
				logStr := ""
				for _, logKey := range tmp {
					if value, ok := mp.Load(logKey); ok {
						// 直接断言为字符串，如果失败，说明数据异常
						if str, ok := value.(interface{}).(string); ok {
							logStr += str + "\n"
						} else {
							fmt.Printf("Unexpected value type for key %s: %T\n", logKey, value)
							// 可以添加更详细的错误处理，比如记录到错误日志
						}
						mp.Delete(logKey)
					}
				}
				l.SizeCheck()
				fmt.Fprintln(l.fileHandler, logStr)
			}
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
