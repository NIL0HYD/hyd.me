package slice

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	myinfo := map[string]string{"name": "huangyuandong", "email": "yuandong.huang@gmail"}

	for v, k := range myinfo {
		fmt.Printf("%v: %v\n", v, k)
	}

	fmt.Println("--------------")

}

func TestMap2(t *testing.T) {
	intmap := map[int]int{
		1: 10, 2: 20,
	}
	for v, k := range intmap {
		fmt.Printf("%v: %v\n", v, k)
	}

	t.Fail()
}
