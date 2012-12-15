package main

import (
	"fmt"
	"math"
)

func Dplus(x int) int

func dplus(x int) int {
	return x + 2
}

func main() {
	fmt.Println(math.Dim(5.0, 4.0))

	fmt.Println(math.Pow(3, 2))

	fmt.Println(math.Pow10(2))

	fmt.Println(Dplus(5))
}
