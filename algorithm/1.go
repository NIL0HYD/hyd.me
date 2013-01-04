package main

import (
	"fmt"
)

/*
【程序1】  题目:古典问题:有一对兔子,从出生后第3个月起每个月都生一对兔子,小兔子长到第四个月后每个月又生一对兔子,假如兔子都不
死,问每个月的兔子总数为多少?  
1.程序分析:   兔子的规律为数列1,1,2,3,5,8,13,21....


*/

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d ", f(i))
	}
}

func f(month int) int {
	if month == 1 || month == 2 {
		return 1
	}
	return f(month-1) + f(month-2)
}
