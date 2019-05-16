package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	command := exec.Command("bash", "getVersion.sh", "arcee")
	//command := exec.Command("uname","-a")
	bytes, err := command.Output()
	if err != nil {
		fmt.Println("err", err)
	}
	result := string(bytes)
	tmp := strings.Split(result, "\n")
	fmt.Println(tmp)

}
