package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetReqInfo(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func PostReqInfo() {
	resp, err := http.Post("http://localhost:8080/post", "application/json", strings.NewReader("{\"name\":\"test\"}"))
	if err != nil {
		fmt.Println("postreq info 报错了: ", resp.Body)
	}
	content, reErr := ioutil.ReadAll(resp.Body)
	if reErr == nil {
		fmt.Println("postreqInfo 成功了: ", string(content))
	}
}

func PostReqInfoWithHeader() {
	data := url.Values{"start": {"0"}, "offset": {"1"}}
	body := strings.NewReader(data.Encode())
	req, reqErr := http.NewRequest("POST", "http://localhost:8080/form", body)
	if reqErr != nil {
		fmt.Println("postwithheader 报错了！")
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	clt := http.Client{}
	resp, respErr := clt.Do(req)
	if respErr == nil {
		content, readErr := ioutil.ReadAll(resp.Body)
		if readErr == nil {
			fmt.Println("postwithheader info : ", string(content))
		}
	}
}

func GetInfo() {
	resp, err := GetReqInfo("http://localhost:8080/test")
	if err != nil {
		fmt.Println("获取数据报错了")
		return
	}
	content, readErr := ioutil.ReadAll(resp.Body)
	if readErr == nil {
		fmt.Println("getinfo返回结果: ", string(content))
	}
}

func main() {
	GetInfo()
	PostReqInfo()
	PostReqInfoWithHeader()
}
