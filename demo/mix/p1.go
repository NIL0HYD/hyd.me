package main

import (
	"fmt"
)

// 1. 当defer调用函数的时候, 函数用到的每个参数和变量的值也会被计算
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 2. Defer调用的函数将在当前函数返回的时候, 以后进先出的顺序执行.
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// 3. Defer调用的函数可以在返回语句执行后读取或修改命名的返回值.
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()
	fmt.Println("--------------")
	b()
	fmt.Println("--------------")
	fmt.Println(c())
}
