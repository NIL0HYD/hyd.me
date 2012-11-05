package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			ch <- 0
		}()
	}

	for i := 0; i < 5; i++ {
		<-ch
	}
}
