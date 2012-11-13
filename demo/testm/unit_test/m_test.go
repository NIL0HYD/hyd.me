package testm

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	var a int
	fmt.Println("a: ", a)
	fmt.Println("&a: ", &a)
	var p *int
	fmt.Println("p: ", p)
	fmt.Println("&p: ", &p)
	p = &a
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)
	fmt.Println("--------------")

	var s string
	fmt.Println("s: ", s)
	if s == "" {
		fmt.Println("s is ")
	}
	fmt.Println("&s: ", &s)
	t.Fail()
}
