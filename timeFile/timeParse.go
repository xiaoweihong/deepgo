package main

import (
	"fmt"
	"strings"
	"time"
)

var shortMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

func main() {
	time_a := "06-Feb-2018"
	timeLayout := "2006-01-02"
	string_tmp := strings.Split(time_a, "-")
	year := string_tmp[2]
	month := string_tmp[1]
	day := string_tmp[0]
	switch month {
	case "Jan":
		month = "01"
	case "Feb":
		month = "02"
	case "Mar":
		month = "03"
	case "Apr":
		month = "04"
	case "May":
		month = "05"
	case "Jun":
		month = "06"
	case "Jul":
		month = "07"
	case "Aug":
		month = "08"
	case "Sep":
		month = "09"
	case "Oct":
		month = "10"
	case "Nov":
		month = "11"
	case "Dec":
		month = "12"
	}
	parse_string := fmt.Sprintf("%s-%s-%s", year, month, day)
	fmt.Println(parse_string)
	fmt.Println(time.Parse(timeLayout, parse_string))
}
