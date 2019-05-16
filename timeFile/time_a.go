package main

import (
	"fmt"
	"github.com/noaway/dateparse"
)

func main() {
	//timeString:="23-Oct-2017"
	a := "12 Feb 2006,"

	parse, err := dateparse.ParseAny(a)
	//timeLayout := "2006-01-02"
	//parse, err := time.Parse(timeLayout, timeString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)
}
