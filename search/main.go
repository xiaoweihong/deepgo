package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	//log.Flags()
	array := [...]int{1, 2, 3, 4, 5}
	array2 := [5]int{12, 3, 3, 3}
	array3 := [...]*int{new(int), new(int)}
	log.Println(array)
	log.Println(array2)
	a := 100
	*array3[0] = a
	log.Println(array3)
	s := make([]int, 3, 77)
	log.Println(s[2])
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 9)
	log.Println(newSlice)
	log.Println(slice)
	//log.Println(cap(newSlice))
}
