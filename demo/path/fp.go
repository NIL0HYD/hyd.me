package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println(filepath.Abs("a/b/c"))
	fmt.Println(filepath.Abs("/"))

	fmt.Println(filepath.EvalSymlinks("F:\\Google\\gopath\\src\\hyd.me\\path\\p.go"))

	fmt.Println(filepath.FromSlash("a/b///c"))

	fmt.Println(os.PathSeparator)
}
