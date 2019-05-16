package main

import (
	"fmt"
	"time"
)

func main() {
	s := "2018-02-02"
	s1 := "2018-03"
	s2 := "2018"
	s3 := "2014-01-08 09:04:41"

	fmt.Println(len(s))
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	t, _ := time.Parse("2006-01-02 15:04:05", s3)
	fmt.Println(t)

}
