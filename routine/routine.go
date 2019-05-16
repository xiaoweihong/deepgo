package main

import (
	"fmt"
	"time"
)

var complete = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now(), i)
	}
	complete <- 0
}

func main() {
	go loop()
	<-complete
	//go loop()
	//<-complete
}
