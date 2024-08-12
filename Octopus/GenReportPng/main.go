package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type config struct {
	JsFile     string `json:"js_file"`
	SaveDir    string `json:"save_dir"`
	RequestURL string `json:"request_url"`
	NodePath   string `json:"node_path"`
}

type rsp interface {
	success(w http.ResponseWriter, data interface{})
	fail(w http.ResponseWriter, data interface{})
}

type resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (rcv *resp) success(w http.ResponseWriter, data interface{}) {
	rcv.Code = 400200
	if rcv.Msg == "" {
		rcv.Msg = "success"
	}

	rcv.content(w, data)
}

func (rcv *resp) fail(w http.ResponseWriter, data interface{}) {
	rcv.Code = 400400
	if rcv.Msg == "" {
		rcv.Msg = "fail"
	}

	rcv.content(w, data)
}

func (rcv *resp) content(w http.ResponseWriter, data interface{}) {
	// 设置响应的内容类型为application/json
	rcv.Data = data
	w.Header().Set("Content-Type", "application/json")
	marshal, err := json.Marshal(rcv)
	if err != nil {
		return
	}

	_, err = w.Write(marshal)
	if err != nil {
		return
	}
}

var conf config

func init() {
	// 解析地址
	parseConfig()
	log.Println("init ok...")

}

func main() {
	http.HandleFunc("/genReport", genHandle)
	log.Println("serve at:", "127.0.0.1:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func parseConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Println(err)
	}

	return
}

func genHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("start...")
	response := &resp{}
	query := r.URL.Query()
	log.Println(r.URL)

	if !query.Has("id") {
		log.Println("missing id")
		response.Msg = "missing id"
		response.fail(w, "")
		return
	}

	//command := conf.command
	// TODO 解析脚本   "C:\Program Files\nodejs\node.exe" E:\phpEnv\www\ct1_vue_jlpjw\public/puppeteer\js\build_file.js http://192.168.1.88:259/#/publicReport?id= E:\phpEnv\www\ct1_vue_jlpjw\public/puppeteer\png\ 229 png
	nodePath := conf.NodePath
	args := []string{
		conf.JsFile,     // TODO project js file path
		conf.RequestURL, // TODO request_path
		conf.SaveDir,    // TODO project file path
		query.Get("id"), // TODO report id 这个就读取
		`png`,           // TODO file type 这个可以暂时不用管
	}

	output, err := exec.Command(nodePath, args...).Output()
	if err != nil {
		log.Println("command run error:", err)
		response.Msg = "command run error"
		response.fail(w, "")
		return
	}

	log.Println(string(output))

	response.success(w, nil)
	log.Println("finish...")
}

func clearCommand() {
	var command string
	if runtime.GOOS == "windows" {
		command = "cls"
	} else {
		command = "clear"
	}

	_, err := exec.Command("cmd", command).Output()
	if err != nil {
		log.Println("clear screen with error:", err)
	}
}
