package main

import (
	"flag"
	"fmt"
)

var (
	name  = flag.String("myname", "huangyuandong", "Nothing")
	email = flag.String("myemail", "yuandong.huang@gmail.com", "email")
	qq    = "176458203"
)

func main() {
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*email)
	fmt.Println(qq)
}
