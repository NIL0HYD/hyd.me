package main

import (
	"fmt"
)

func go1(emailSuffix string) {
	c <- "yuandong.huang@" + emailSuffix
}

func main() {
	c := make(chan string)
	go go1("gmail.com")
	go go1("163.com")

	fmt.Printf("main: %s \n", <-c)
	fmt.Printf("main: %s \n", <-c)
}
