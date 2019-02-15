package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("./config.json")

	if err != nil {

	}
	//var str_byte []byte
	map2 := make(map[string]interface{})
	s, err := ioutil.ReadAll(f)
	fmt.Println(s)
	json.Unmarshal(s, &map2)

	fmt.Println(map2["mongo"].(map[string]interface{})["a"])
}
