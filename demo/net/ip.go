package main

import (
	"log"
	"net"
)

func main() {
	addr := net.ParseIP("10.10.71.129")
	log.Print(addr.String())
}
