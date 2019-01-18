package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.2.66:6379")
	if err != nil {
		panic("error")
		os.Exit(-1)
	}
	defer conn.Close()
	//_,err=conn.Do("set","name","xiaowei2")
	//
	//if err != nil {
	//
	//}
	//name,err := redis.String(conn.Do("get","name"))
	//fmt.Println(name)

	a, err := conn.Do("select", "1")
	fmt.Println(a)
	result, err := redis.StringMap(conn.Do("hgetall", "supmylo:sys:stats"))
	fmt.Println(result["192.168.2.66"][1:22])

}
