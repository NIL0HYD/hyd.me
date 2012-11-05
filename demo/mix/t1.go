package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	time.Sleep(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
