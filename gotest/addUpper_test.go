package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(11)

	if res != 55 {
		//fmt.Println("AddUpper(10) error,期望%v,")
		t.Fatalf("AddUpper(10) error,期望%v,实际%v", 55, res)
	}
	t.Logf("AddUpper(10) 执行正确")
}
