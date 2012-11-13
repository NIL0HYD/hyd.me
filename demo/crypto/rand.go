package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	if n, err := rand.Int(rand.Reader, big.NewInt(1024)); err == nil {
		fmt.Println(n)
	}

	if n, err := rand.Prime(rand.Reader, 128); err == nil {
		fmt.Println(n)
	}

	//25885087847758063
	//172947142856842446121302980611922498603

	d := 0x12
	fmt.Printf("%d", d)
}
