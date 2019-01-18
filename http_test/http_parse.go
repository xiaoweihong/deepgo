package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type result struct {
	Message string
}

func main() {

	//var s result
	//json.NewDecoder(resp.Body).Decode(&s)
	//fmt.Println(s)

}

func getUrlResponse(url string) (resp string, err error) {
	resp2, err := http.Get(url)
	if err != nil {
	}
	defer resp2.Body.Close()

	body, _ := ioutil.ReadAll(resp2.Body)

	fmt.Println(string(body))
	return string(body), err
}
