package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	//t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("type: ", reflect.TypeOf(x))
	fmt.Println("Value: ", reflect.ValueOf(x))
	fmt.Println("----")
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())
}
