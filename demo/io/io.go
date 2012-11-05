package main

import (
	"encoding/json"
	"fmt"
)

type BtStru struct {
	Name string
	Age  int32
}

func New() *BtStru {
	return &BtStru{}
}

type slicer []byte

func (sp *slicer) next(n int) (b []byte) {
	s := *sp
	b, *sp = s[0:n], s[n:]
	return
}

func main() {
	bs := New()
	bs.Name = "中文呢"
	bs.Age = 26
	fmt.Println([]byte(bs.Name))
	fmt.Println("---------------------")
	jsonStyle(bs)
}

func jsonStyle(bs *BtStru) {
	bys, _ := json.Marshal(bs)
	fmt.Println(string(bys))
	fmt.Println(bys)
	fmt.Println("---------------")
	sp := slicer(bys)
	for i := 0; i < len(sp); i++ {
		fmt.Println("len(sp): ", len(sp))
		if len(sp) != 0 {
			fmt.Println(sp.next(3))
		}
	}
}
