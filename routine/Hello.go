package main

import (
	"fmt"
	"time"
)

func sample1(message chan string) {
	message <- "Hello world.1"
	message <- "Hello world.2"
	message <- "Hello world.3"
	message <- "Hello world.4"

}

func sample2(message chan string) {
	//time.Sleep(2*time.Second)
	str := <-message
	str = str + "I'm go routine"
	message <- str
	close(message)
}

func main() {

	var message = make(chan string, 3)
	go sample1(message)
	go sample2(message)
	time.Sleep(1 * time.Second)
	for str := range message {
		fmt.Println(str)
	}

}
