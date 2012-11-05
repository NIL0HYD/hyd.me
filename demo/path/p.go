package main

import (
	"fmt"
	"path"
)

func main() {

	fmt.Println(path.Join("/", "a", "b", "c"))

	fmt.Println()
	fmt.Println(path.IsAbs("/dev/null"))

	fmt.Println()
	p1 := "/a/b"
	fmt.Println(path.Base(p1)) // output: b

	p2 := ""
	fmt.Println(path.Base(p2)) // output: .

	p3 := "/"
	fmt.Println(path.Base(p3)) // output: /

	p4 := "/a/b/c/"
	fmt.Println(path.Base(p4)) // output: c

	fmt.Println()

	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

	fmt.Println()
	fmt.Println(path.Dir("/a/b/c"))

	fmt.Println()
	fmt.Println(path.Ext("/a/b/c/bar.css"))

	fmt.Println()
	fmt.Println(path.Split("/a/static/myfile.css"))
	fmt.Println(path.Split("/myfile.css"))

}
