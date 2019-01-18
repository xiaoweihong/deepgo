package main

import "fmt"

func main() {
	a := 1
	var b *int
	b = &a
	fmt.Println(a, &a, *b)
}
