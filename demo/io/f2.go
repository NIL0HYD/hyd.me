package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := http.Get("http://127.0.0.1:6060/pkg/net/http/")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Printf("%s\n", string(b))
	}
}
