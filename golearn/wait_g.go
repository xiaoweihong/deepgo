package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func add(a, b int) {
	defer wg.Done()
	c := a + b

	fmt.Println(c)
}
func main() {
	wg.Add(20)
	go func() {

		defer wg.Done()
		fmt.Println(1, 2)
	}()

	go add(1, 2)

	wg.Wait()
}
