package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)
	for i := 0; i < 5; i++ {
		// go starts a goroutine
		go printHelloWorld(i, ch) // go调用函数为并发执行
	}

	// 主程序拿到返回来的ch输出
	for i := 0; i < 5; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}

func printHelloWorld(i int, ch chan string) {
	for {
		// 把print中的东西传递给ch
		ch <- fmt.Sprintf("HelloWorld from goroutine %d!\n",
			i)
	}
}
