package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	sl := make([]int, 10)
	for v, k := range sl {
		fmt.Printf("%v, %v\n", v, k)
	}

	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	fmt.Println("--------------------")
}

func TestArraySlice(t *testing.T) {
	array := [...]int{1, 2, 3, 4, 5}
	sl := array[1:4]

	for v, k := range array {
		fmt.Printf("%v, %v\n", v, k)
	}
	fmt.Println(len(array))
	fmt.Println(cap(array))

	fmt.Println("-------------")
	for v, k := range sl {
		fmt.Printf("%v, %v\n", v, k)
	}

	fmt.Println(len(sl))
	fmt.Println(cap(sl))

}

func TestAppend(t *testing.T) {
	fmt.Println("----------------")
	sl := []int{1, 2, 3}
	fmt.Println(&sl)
	sl = append(sl, 4)
	fmt.Println(&sl)
	fmt.Println(sl[3])
	fmt.Println("----------------")
}

func TestYw(t *testing.T) {
	fmt.Println("----------------")
	var i uint
	i = 20
	fmt.Println(1 << i)
	fmt.Println("----------------")
	t.Fail()
}
