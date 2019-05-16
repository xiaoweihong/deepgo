package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := "http://192.168.2.25:9333/11,2ce6a06c12b3"

	//fmt.Println(strings.Replace(s,"8501/api/file","9333",-1))

	resp, err := http.Head(s)
	if err != nil {

	}
	fmt.Printf("%T", resp.Request.URL.String())
}
