package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

var cookieMap = make(map[string]string)

// TODO construct answer
func generateAnswer(evl evaluates, m *[]questions, c chan *[]byte) {
	py := payload{
		UserId:    evl.UserId,
		MeasureId: evl.MeasureId,
		IsSave:    true,
		Access:    access{accessToken: "md5"},
		Inputs:    struct{}{},
		Answer:    nil,
		Id:        evl.Id,
	}

	var answers []answer
	for _, value := range *m {

		answers = append(answers, answer{
			Id:     value.Id,
			Answer: getAnswer(value.Score),
		})
	}

	py.Answer = answers
	x, err := json.Marshal(py)
	if err != nil {
		log.Println("generateAnswer: json marshal with error: ", err)
	}
	c <- &x
}

// TODO if configure MAX_SCORE or MIN_SCORE ,get the score's choice, if not,get random choice
func getAnswer(score string) string {
	s := strings.Split(score, " ")
	x := rand.Intn(len(s))
	return string(rune(x + 'A'))
}

func login() string {
	loginUrl := fmt.Sprintf("%s:%d%s", cnf.Require.Uri, cnf.Require.Port, cnf.Require.LoginUri)
	data := []byte(fmt.Sprintf(`{"userName": %s, "password": %s,"type":2}`, "admin", "123456"))
	//if cookieStr, ok := cookieMap[string(data)]; ok {
	//	return cookieStr
	//}

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

	cookieMap[string(data)] = cookieStr

	return cookieStr
}

// req send a request. return *http.Response, you must close response.Body while use it over
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

// TODO commit with cookie
func commit(evl evaluates, m *[]questions, ch chan *http.Response) {

	var wg sync.WaitGroup
	c := make(chan *[]byte, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		generateAnswer(evl, m, c)
	}()

	wg.Wait()
	payloads := <-c

	cookie := login()

	commitUrl := fmt.Sprintf("%s:%d%s", cnf.Require.Uri, cnf.Require.Port, cnf.Require.CommitUri)
	go func() {
		response, err := req(*payloads, commitUrl, "POST", cookie)
		if err != nil {
			log.Println("go func req with error:", err)
		}
		log.Println(string(*payloads), commitUrl, cookie)
		ch <- response
	}()

}
