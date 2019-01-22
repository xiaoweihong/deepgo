package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getNameLen(name string) int {
	return len(name)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message + " : " + fmt.Sprintf("%d", getNameLen(message))
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
