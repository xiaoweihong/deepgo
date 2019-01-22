package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://baidu.com"
	fmt.Println("URL:>", url)
	//
	////登陆用户名
	usrId := "LaoWong"
	//
	////登陆密码
	pwd := "pwd1234"

	//json序列化
	post := "{\"UserId\":\"" + usrId +
		"\",\"Password\":\"" + pwd +
		"\"}"

	fmt.Println(url, "post", post)

	var jsonStr = []byte(post)
	fmt.Println("jsonStr", jsonStr)
	fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
