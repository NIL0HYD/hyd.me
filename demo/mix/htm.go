package main

import (
	"fmt"
	"html"
)

func main() {
	ht := "<h1>huanghuanghuang</h1>"

	fmt.Println(html.EscapeString(ht))

	uht := "&lt;h1&gt;huanghuanghuang&lt;/h1&gt"
	fmt.Println(html.UnescapeString(uht))
}
