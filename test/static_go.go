//go:generate statik -src=./assets
//go:generate go fmt statik/statik.go

package main

import (
	_ "deepgo/test/statik"
	"fmt"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"os"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	file, err := statikFS.Open("/abc")
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}
