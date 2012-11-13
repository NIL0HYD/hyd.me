package main

import (
	"fmt"
)

func main() {
	ar := [...]string{0: "no error", 1: "invalid argument"}
	for k, v := range ar {
		fmt.Println(k, v)
	}

	fmt.Printf("%f", 1e9) //1000000000.000000
}
