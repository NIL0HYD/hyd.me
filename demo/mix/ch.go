package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Printf("%d, ", <-ch)
	ch <- 3
	fmt.Printf("%d, %d", <-ch, <-ch)
}
