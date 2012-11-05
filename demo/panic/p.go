package main

import (
	"errors"
	"log"
)

var ErrFlase = errors.New("This is a error!")

func throw() {
	panic(ErrFlase)
}

func second() {
	defer func() {
		log.Print("Second defer")
		err := recover()
		log.Print(err)
	}()
	throw()
}

func third() {
	defer func() {
		log.Print("Third defer")
	}()
	second()
}

func main() {
	defer func() {
		str := recover()
		log.Print(str)
		log.Print("Any way it run")
	}()
	log.Print("Thinging about something!")
	third()
	log.Print("main")
}
