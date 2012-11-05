package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {
	s := "this is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
	fmt.Println("-------------------")

	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(reflect.ValueOf(x).Interface())
	fmt.Println("-------------------")

	a := new(MyStruct)
	a.name = "nnnnnnnnnnnnnnnnnnnnnnnnnnnn"
	typ := reflect.TypeOf(a)

	fmt.Println(typ.NumMethod())
	fmt.Println("-------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])

}
