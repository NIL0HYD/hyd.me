package main

import (
	"fmt"
)

/*
&    bitwise and 按位与           integers
|    bitwise or  按位或           integers
^    bitwise xor 按位异或           integers
&^   bit clear (and not)    integers
*/
const (
	GREAT = 9
	NICE  = 5
	OR    = GREAT | NICE
	AND   = GREAT & NICE

	ODD  = 3
	EVEN = 8
)

func main() {
	fmt.Printf("%b | %b | %b:%v \n", GREAT, NICE, OR, OR)

	fmt.Printf("%b | %b | %b:%v \n", GREAT, NICE, AND, AND)

	fmt.Println(ODD&1 == 1)
	fmt.Println(EVEN&1 == 1)

	// 交换两个整数，不引入第三个变量
	a, b := 5, 3
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a)
	fmt.Println(b)
}
