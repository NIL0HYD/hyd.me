package main

import (
	"fmt"
	"index/suffixarray"
)

// golang.org网站的全文搜索是基于suffix array实现的【http://t.cn/hBJekg】，
// 可能觉得效果不错，就把suffix array添加到golang的标准库里面了。【http://t.cn/hBJekd】
// http://blog.csdn.net/fxsjy/article/details/6297523
func main() {
	fmt.Println("Hello, 世界")
	str := `The Go programming language is an open source project to make programmers more productive. 
	Go is expressive, concise, clean, and efficient. 
	Its concurrency mechanisms make it easy to write programs that get the most out of multicore 
	and networked machines, hyd, while its novel type system enables flexible 
	and modular program construction. 
	Go compiles quickly to machine code yet has the convenience of garbage collection 
	and the power of run-time reflection. It's a fast, statically typed, 
	compiled language that feels like a dynamically typed, interpreted language.`

	index := suffixarray.New([]byte(str))
	offsets1 := index.Lookup([]byte("hyd"), -1)
	for _, i := range offsets1 {
		fmt.Println(str[i:])
	}

}
