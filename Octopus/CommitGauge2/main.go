package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"os"
	"sync"
)

// mysql 数据库配置
type mysqlCnf struct {
	User      string `json:"user"`      // JSON键 "user" 映射到此字段
	Password  string `json:"password"`  // JSON键 "password" 映射到此字段
	Host      string `json:"host"`      // JSON键 "host" 映射到此字段
	Port      string `json:"port"`      // JSON键 "port" 映射到此字段
	DbName    string `json:"dbname"`    // JSON键 "dbname" 映射到此字段
	CycleTime int    `json:"cycleTime"` // JSON键 "cycleTime" 映射到此字段
}

type requireCnf struct {
	Uri       string `json:"domain"`     // 请求的的地址
	Port      int    `json:"port"`       // 请求端口
	LoginUri  string `json:"login_uri"`  // 登录地址
	CommitUri string `json:"commit_uri"` // 提交地址
}

type config struct {
	Mysql   mysqlCnf   `json:"mysql"`
	Require requireCnf `json:"require"`
	User    []int      `json:"user"`
	Gauge   []int      `json:"gauge"`
}

// 提交量表的数据结构
type payload struct {
	UserId    int      `json:"user_id"`
	MeasureId int      `json:"measureId"`
	IsSave    bool     `json:"isSave"`
	Access    access   `json:"systemAccess"`
	Inputs    struct{} `json:"inputs"`
	Answer    []answer `json:"answer"`
	Id        int      `json:"id"`
}

type access struct {
	accessToken string
}

type answer struct {
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

type evaluates struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	MeasureId int `json:"measure_id"`
}

var (
	maxWorkers = 2 // 最大并发数
	sema       = semaphore.NewWeighted(int64(maxWorkers))
	wg         sync.WaitGroup
)

var index = 0

var lock = sync.RWMutex{}

func increment() int {
	lock.Lock()
	defer lock.Unlock() // 确保在函数退出时释放锁，防止死锁
	tmp := index
	index++
	return tmp
}

type questions struct {
	Id      int    `json:"id"`
	Score   string `json:"score"`
	Options string `json:"options"`
}

type measureMap map[int][]questions

var cnf *config

func init() {
	cnf = parseConfig()
}

func main() {
	// parse config and save
	//config := parseConfig()
	//fmt.Printf("%#v\n", cnf)
	// == 上面的user_config 用来筛选用户罢了
	// 上面的gauge_config 也是用来筛选量表罢了。统一定位到
	// 1. 查询出来有所得evaluate_log 只要id，measure_id,user_id就行了
	// 获取到了待做的报告，以及量表，下一步就是构造题目了。
	evl := getEvaluate()
	fmt.Printf("%#v", evl)
	me := getMeasures(evl)
	fmt.Printf("%#v", me)
	for k, v := range *me {
		fmt.Printf("%s %d", k, len(v))
	}
}

func parseConfig() *config {
	var config *config
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Println("open file with err:", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("unmarshal json config with error:", err)
	}

	return config
}
