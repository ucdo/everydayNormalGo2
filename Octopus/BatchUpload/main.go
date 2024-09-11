package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type conf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
	Target   string `json:"dir"`
}

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"data"`
}

var uri = map[string]string{}

var cnf *conf

var cookie string

func init() {
	cnf = getConfig()
	uri["login"] = "/login"
	uri["upload"] = "/file_upload"
	uri["create"] = "/created_measure"
	login()
}

func main() {
	scan()
	os.Exit(1)
}

func scan() {
	var wg sync.WaitGroup
	entries, err := os.ReadDir(cnf.Target)
	if err != nil {
		log.Panicln(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // 忽略目录
		}
		//log.Println(entry.Name())

		filePath := filepath.Join(cnf.Target, entry.Name())
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			upload(filePath)
		}(filePath)
	}
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("All tasks are completed.")

}

func getConfig() *conf {
	file := "config.json"
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cnf := &conf{}
	err = json.Unmarshal(data, cnf)
	if err != nil {
		panic(err)
	}

	return cnf
}

func login() {
	loginUrl := fmt.Sprintf("%s%s", cnf.Domain, uri["login"])
	data := []byte(fmt.Sprintf(`{"userName": "%s", "password": "%s","type":"%d"}`, cnf.User, cnf.Password, 2))

	response, err := req(data, loginUrl, "POST", "")
	if err != nil {
		log.Println("login with err: ", err, ". login data:", string(data))
	}

	var cookieStr string
	for i, cookie := range response.Cookies() {
		if i > 0 {
			cookieStr += "; "
		}
		cookieStr += cookie.Name + "=" + cookie.Value
	}
	e, _ := io.ReadAll(response.Body)
	log.Printf("login result:%s\n", e)

	cookie = cookieStr
}

func req(data []byte, url string, method string, cookie string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Host", url)
	request.Header.Add("Connection", "keep-alive")
	if len(cookie) > 0 {
		request.Header.Add("Cookie", cookie)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func uploadReq(filePath string, url string) (*http.Response, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建一个缓冲，用于存储表单数据
	var requestBody bytes.Buffer

	// 创建一个multipart writer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件部分
	fileWriter, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = fileWriter.Write(readFile(file))
	if err != nil {
		return nil, err
	}

	// 添加其他表单字段
	err = writer.WriteField("type", "4")
	if err != nil {
		return nil, err
	}
	err = writer.WriteField("attribution_id", "1")
	if err != nil {
		return nil, err
	}

	// 关闭writer，这将写入请求体的结束标志
	writer.Close()

	// 创建请求
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return nil, err
	}

	// 设置Content-Type为multipart/form-data，并且boundary是writer生成的
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	request.Header.Add("Cookie", cookie)

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func readFile(file *os.File) []byte {
	e, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return e
}

func upload(path string) {

	response, err := uploadReq(path, fmt.Sprintf("%s%s", cnf.Domain, uri["upload"]))
	if err != nil {
		log.Println("login with err: ", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	//log.Println(string(body))

	response.Body.Close()
	e := &ResponseData{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		log.Println("json unmarshal err:", err)
	}

	//log.Printf("%#v\n", e.Data.URL)

	create(e.Data.URL, path)
}

func create(respUrl string, path string) {
	url := fmt.Sprintf("%s%s", cnf.Domain, uri["create"])
	data := []byte(fmt.Sprintf(`{"classifyId": "%d", "customDesc": "","customName":"","fileList":[],"measurePath":"%s","theme":"0"}`, 87, respUrl))

	response, err := req(data, url, "POST", cookie)
	if err != nil {
		log.Println("login with err: ", err, ". login data:", string(data))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("io.ReadAll(response.Body) err:", err)
	}

	body2 := string(body)
	if strings.LastIndex(body2, "添加成功！") == -1 {
		log.Printf("create reuslt with err: %s", path)
	}

}
