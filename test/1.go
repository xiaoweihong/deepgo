package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a    bc"
	fields := strings.Fields(s)
	fmt.Println(fields)
	fmt.Println(len(fields))

}
