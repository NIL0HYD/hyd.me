package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	util()
}

func bcopy() {
	rb := []byte("HuanYuanDong")
	wb := make([]byte, len(rb))
	copy(wb, rb) // 果然是buildin的函数
	fmt.Println(string(wb))
}

func util() {
	var dir = "."
	if fis, err := ioutil.ReadDir(dir); err == nil {
		for _, fi := range fis {
			fmt.Println(fi.Name())
		}
	}
}
