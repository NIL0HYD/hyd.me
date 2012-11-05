package iface

import (
	"fmt"
)

type Sa struct {
	a int
}

func (sa *Sa) Get() int {
	return sa.a
}

func (sa *Sa) Put(v int) {
	sa.a = v
}

func (sa Sa) say() {
	fmt.Println("what is this?")
}
