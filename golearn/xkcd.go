package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const xkcdApi = "https://xkcd.com/"

type xkcd_result struct {
	Month     string
	Num       int
	Link      string
	Year      string
	News      string
	SafeTitle string `json:"safe_title"`
	Img       string
}

func getResult(url string, ch chan xkcd_result) {

	urlfull := xkcdApi + url + "/info.0.json"
	//fmt.Println(urlfull)
	resp, err := http.Get(urlfull)
	if err != nil {
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
	}
	var result xkcd_result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
	}
	resp.Body.Close()
	ch <- result
}

func main() {
	ch := make(chan xkcd_result)
	for i := 10; i < 15; i++ {
		go getResult(strconv.Itoa(i), ch)
	}
	for i := 10; i < 15; i++ {
		msg := <-ch
		fmt.Println(msg)
	}

}
