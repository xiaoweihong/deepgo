package main

import (
	"deepgo/file"
	"fmt"
	"os"
)

func main() {
	s, _ := os.Executable()
	fmt.Println(s)
	fmt.Println(file.SelfDir())
}
