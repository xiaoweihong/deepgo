package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println("a")
	}
	c := os.Stdout
	a, err := io.Copy(c, resp.Body)
	if err != nil {

	}
	fmt.Printf("%s", a)
	//b,err:=ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {

	}
	//fmt.Printf("%s", b)

}
