package main

import "fmt"

// 好奇怪的东西？
type MyFunc func(string)

func (this MyFunc) Get(info string) {
	this(info)
}

func status(s string) {
	fmt.Println(s)
}

func main() {
	a := MyFunc(status)
	a.Get("Oops..")
	a("Oops..")
}
