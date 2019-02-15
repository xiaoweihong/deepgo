package main

import (
	"fmt"
	"strconv"
)

func main() {
	//t:=time.Now().UnixNano()
	//a:=t/1e9
	//fmt.Println(a)
	//fmt.Println(t)
	//
	//fmt.Printf("%T",time.Unix(a,0).Format("2006-01-02 15:04:05"))

	s := "101.1"
	abc, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(abc)

	var a = []int{1, 2, 3, 4}
	fmt.Println(a)
}
