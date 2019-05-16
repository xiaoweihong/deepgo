package utils

import (
	"fmt"
	"net"
	"os"
)

func GetIp() []string {
	addrs, err := net.InterfaceAddrs()
	var ipList []string
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipList = append(ipList, ipnet.IP.String())
				return ipList
			}

		}
	}
	return ipList
}
