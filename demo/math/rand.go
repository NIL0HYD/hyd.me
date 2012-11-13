package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// uses the provided seed value to initialize the generator to a deterministic state. 
	// If Seed is not called, the generator behaves as if seeded by Seed(1).
	rand.Seed(time.Now().Unix())

	n := rand.Int()

	k := rand.Intn(100)

	fmt.Println(n, k)

	// 16进制
	xx := []int{0x00, 0x01, 0x02, 0x03, 0x04, 0x05,
		0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12}
	for _, v := range xx {
		fmt.Println(v)
	}
}
