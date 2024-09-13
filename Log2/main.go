package main

import (
	"Log2/myLog"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type config struct {
	LogLevel string `cnf:"log_level"`
	FileName string `cnf:"file_name"`
	FilePath string `cnf:"file_path"`
	FileSize int    `cnf:"file_size"`
}

var wg sync.WaitGroup

var Logger myLog.MyLogger

func main() {
	parseConf("./main.cnf")
	// call
	//now := time.Now().Format("2006-01-02") + "_1.log"
	//logger := myLog.NewFileLog("debug", "./runtime/", now)
	//
	//for i := 0; i < 1e6; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		logger.Warn("写写warn")
	//		logger.Fatal("用户%4d正在疯狂尝试登录", i)
	//	}(i)
	//}
	//wg.Wait()
}

func parseConf(path string) {
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

	fmt.Printf("%#v", cnf)
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
