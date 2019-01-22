package main

import (
	"flag"
	"github.com/golang/glog"
	"net/http"
)

func init() {
	flag.Set("alsologtostderr", "true")
}

var resp_chanel = make(chan interface{})

func getRequest(url string) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	//s,err:=io.Copy(os.Stdout,resp.Body)
	//resp_chanel<-s
	resp_chanel <- resp.StatusCode

	defer resp.Body.Close()
}

func main() {

	flag.Parse()
	for i := 0; i < 50; i++ {
		go getRequest("")
	}

	for i := 0; i < 50; i++ {
		a := <-resp_chanel
		glog.Infoln(a)
	}

	//for{
	//	go getRequest("")
	//	time.Sleep(time.Second)
	//}
	//
	//for{
	//	a:=<-resp_chanel
	//	glog.Infoln(a)
	//}

}
