package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "vehicle10 rtsp://192.168.2.168/live/vehicle10 22.22 22.00"

	split := strings.Split(s, " ")

	fmt.Println(split)

}
