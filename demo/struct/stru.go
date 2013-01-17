package main

import (
	"fmt"
)

/*
从个性（实现）再到一般性（抽象），这是一个归纳的过程，这符合人类对世界的认识。
http://monnand.me/p/ready-go-1/zhCN/
http://monnand.me/p/ready-go-2/zhCN/
*/

type Duck struct {
	Name string
}

func (d *Duck) Eat() {
	fmt.Println(d.Name, "eating...")
}

type NotDuck struct {
	Duck // 匿名成员，实现组合
	Age  int
}

type Animal interface {
	Eat()
}

func main() {
	d := new(Duck)
	d.Name = "Not"
	d.Eat()

	n := new(NotDuck)
	n.Name = "Still"
	n.Age = 10
	n.Eat()

	// 接口
	var a Animal
	a = d
	a.Eat()
	a = n
	a.Eat()
}
