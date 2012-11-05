package main

import "fmt"

func main() {
	if x, flag := 5, true; flag {
		fmt.Println(x)
	}

	fmt.Printf("%b\n", 176458203)
	fmt.Printf("%x\n", "黄远董")
	fmt.Print(fmt.Sprintf("%b\n", 176458203))
	fmt.Print(fmt.Sprintf("%x\n", "黄远董"))
}
