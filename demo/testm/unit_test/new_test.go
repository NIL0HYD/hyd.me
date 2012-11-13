package testm

import (
	"fmt"
	"testing"
)

type SyncedBuffer struct {
	a int
	s string
}

func TestNew(t *testing.T) {
	// 返回一个指针，指向新分配类型int的初始值
	a := new(int)

	fmt.Println("a: ", a)
	fmt.Println("*a: ", *a)

	t.Fail()
}

func TestStruct(t *testing.T) {
	p := new(SyncedBuffer)
	p.a = 1
	p.s = "huanghuanghaunghuang"
	fmt.Println("new p: ", *p)

	var pp SyncedBuffer
	fmt.Println("var pp.a: ", pp.a)
	fmt.Println("var pp.s: ", pp.s)

	t.Fail()
}
