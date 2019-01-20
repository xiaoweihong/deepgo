package main

import (
	"fmt"
	"time"
)

func producer(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s:%v", header, time.Now())
		//time.Sleep(time.Second)
	}
}

func consumer(channel <-chan string) {
	for {
		messsage := <-channel
		fmt.Println(messsage)
	}
}

func main() {
	channel := make(chan string)
	go producer("dog", channel)
	go producer("cat", channel)
	consumer(channel)
}
