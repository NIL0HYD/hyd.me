package main

import (
	"fmt"
	//"runtime"
	"time"
)

func sheep(i int) {
	for ; ; i += 2 {
		fmt.Println(i)
	}
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	go sheep(1)
	go sheep(2)
	time.Sleep(time.Millisecond)
}
