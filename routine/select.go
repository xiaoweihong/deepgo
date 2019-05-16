package main

import (
	"fmt"
	"strconv"
	"time"
)

func sample11(ch chan string) {
	for i := 0; i < 19; i++ {
		ch <- "select test" + strconv.Itoa(i)
		time.Sleep(time.Second * 1)
	}
}

func sample112(ch chan int) {
	for i := 0; i < 19; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	for i := 0; i < 10; i++ {
		go sample11(ch1)
		go sample112(ch2)
	}
	for {

		select {
		case str, ch1Check := <-ch1:
			if !ch1Check {
				fmt.Println("ch1 failed")
			}
			fmt.Println(str)
		case p, ch2Check := <-ch2:
			if !ch2Check {
				fmt.Println("ch2 failed")
			}
			fmt.Println(p)
		}
	}
	//time.Sleep(60*time.Second)
}
