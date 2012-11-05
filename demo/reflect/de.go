package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 3, 2}

	var c, d int

	e := make(map[int]string)
	f := make(map[int]string)
	e[1] = "abc"
	f[1] = "ab" + "c"

	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(c, d))
	fmt.Println(reflect.DeepEqual(e, f))
}
