package main

import (
	"fmt"
	"os"
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

// each source file can define its own init() function to set up whatever state is required.
// 每个源文件都可以定义它自己的init()方法，设置它所需要的任何状态
func init() {
	if USER == "" {
		USER = "Huang Yuandong"
	}
	if HOME == "" {
		HOME = "/usr/" + USER
	}
	if GOROOT == "" {
		GOROOT = HOME + "/go"
	}
}

func main() {
	fmt.Println(HOME)
	fmt.Println(USER)
	fmt.Println(GOROOT)
}
