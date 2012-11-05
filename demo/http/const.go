package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("DefaultMaxHeaderBytes: ", http.DefaultMaxHeaderBytes)

	log.Print(http.Dir("img/"))

}
