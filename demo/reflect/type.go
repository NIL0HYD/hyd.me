package main

import (
	//"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

type MyInt int

var i int
var j MyInt

var r io.Reader

var inter interface{}

func main() {

	r = os.Stdin
	//r = bufio.NewReader(r)
	//r = new(bytes.Buffer)

	inter = 1

	j = 2

	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(j).Kind())
	fmt.Println(reflect.ValueOf(j))

	fmt.Println(reflect.TypeOf(i))

	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.TypeOf(r).Kind())

	fmt.Println(reflect.TypeOf(inter))
	fmt.Println(reflect.TypeOf(inter).Kind())

	fmt.Println(reflect.ValueOf(inter))
	fmt.Println(&inter)

	fmt.Println(reflect.ValueOf(inter).Interface())

	fmt.Println(reflect.ValueOf(r).Interface())

}

/*	
	TypeOf 返回形参值的实际类型
	（TypeOf returns the reflection Type of the value in the interface{}）
	ValueOf 返回一个初始化为存储在interface i中具体值的新值。？？
	（ValueOf returns a new Value initialized to the concrete value stored in the interface i）
	Value is the reflection interface to a Go value.
	Kind 返回底层的类型

	接口类型的变量存储了两个内容：赋值给变量实际的值和这个值的类型描述。
	更准确的说，值是底层实现了接口的实际数据项目，而类型描述了这个项目完整的类型

	接口不能保存`接口值`。

	在基本的层面上，反射只是一个检查`存储在接口变量中`的`类型`和`值`的算法。

	务必记得反射值需要某些内容的地址来修改它指向的东西。
*/
