package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int
}

func show(i interface{}) {
	switch t := i.(type) {
	case *Person:
		fmt.Println(t)
		tt := reflect.TypeOf(i)
		fmt.Println(reflect.ValueOf(i))
		tag := tt.Elem().Field(0).Tag
		name := tt.Elem().Field(0).Name
		fmt.Println(tag)
		fmt.Println(name)
	}

}

func main() {
	p1 := new(Person)
	show(p1)
}
