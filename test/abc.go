package main

import (
	"flag"
	"github.com/golang/glog"
)

func init() {
	flag.Set("logdir", "./")
	flag.Set("alsologtostderr", "true")
	flag.Parse()

}

func main() {
	glog.Info("2123123")
}
