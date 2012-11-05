package main

import (
	"fmt"
)

func sender(ch chan int) {
	//defer func() { recover() }()
	fmt.Println("now i: 00")
	ch <- 1
	fmt.Println("now i: ")
	ch <- 2
	fmt.Println("now i: ")
	ch <- 3
	fmt.Println("now i: ")
}

func reciever(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println("..1.")
	fmt.Println(<-ch)
	fmt.Println("..2.")
	fmt.Println(<-ch)
	fmt.Println("..3.")

}

func main() {

	ch := make(chan int)

	go sender(ch)

	reciever(ch)

	fmt.Println("done!")
}
