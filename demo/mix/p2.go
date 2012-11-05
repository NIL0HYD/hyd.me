package main

import (
	"fmt"
)

func pa() {
	panic("huangyuandong")
}

func main() {

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Name:", v)
		}
	}()

	pa()
	// panic后被忽略
	fmt.Println("Before")

	// panic后被忽略
	fmt.Println("After")
}
