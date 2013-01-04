package main

import (
	"fmt"
)

/*
【程序3】   题目:打印出所有的 "水仙花数 ",所谓 "水仙花数 "是指一个三位数,其各位数字立方和等于该数本身。例如:153是一个 "水仙花
数 ",因为153=1的三次方+5的三次方+3的三次方。  
1.程序分析:利用for循环控制100-999个数,每个数分解出个位,十位,百位。

*/

func main() {
	for i := 100; i < 999; i++ {
		if sxh(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func sxh(x int) bool {
	b := x / 100
	s := (x % 100) / 10
	g := x % 10
	if x == b*b*b+s*s*s+g*g*g {
		return true
	}
	return false
}
