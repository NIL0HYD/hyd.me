package iface

import (
	"fmt"
	"testing"
)

func callf(p I) {
	switch p.(type) {
	case *Sa:
		fmt.Println("*Sa")
	case *Sb:
		fmt.Println("*Sb")
	case Sc:
		fmt.Println("Sc")
	// case Sb:
	// 	fmt.Println("Sb")
	default:
		fmt.Println("defualt")
	}
}

func callff(p I) {

	if t, ok := p.(I); ok {
		fmt.Println(t.(I))
	}

}

func TestSt(t *testing.T) {
	var sa Sa
	callf(&sa)
	callff(&sa)
	sa.say()

	var sc Sc
	callf(sc)

	t.Fail()
}
