package main

import (
	"container/ring"
	"fmt"
)

func printe(s interface{}) {
	fmt.Println(s)
}

func main() {
	r := ring.New(5)
	r.Value = 1
	r.Next().Value = 2
	r.Prev().Value = 5

	plus := ring.New(1)
	plus.Value = 10
	r.Link(plus)
	fmt.Println("len: ", r.Len())
	r.Do(printe)
}
