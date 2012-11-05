package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- true
}

func shower(c chan int, quit chan bool) {
	fmt.Println("-------------------")
	for { // channel的值被修改一次，for运行一次
		fmt.Println("for part--------------")
		select {
		case j := <-c:
			fmt.Printf("%d \n", j)
		case <-quit:
			break
		}
	}
}
