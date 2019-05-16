package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()
	sig := WaitForSignal(syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	fmt.Printf("Gofse exit,got signal: %v, trigger to stop system", sig)
}

func WaitForSignal(sources ...os.Signal) os.Signal {
	var s = make(chan os.Signal, 1)
	defer signal.Stop(s) //the second Ctrl+C will force shutdown

	signal.Notify(s, sources...)
	return <-s //blocked
}
