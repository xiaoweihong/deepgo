package main

import (
	"fmt"
	"github.com/micro/go-config"
)

func main() {
	config.LoadFile("config.json")
	fmt.Println(config.Map()["mongo"])
}
