package main

import (
	//"flag"
	"fmt"
	"os"
	"path/filepath"
)

func getFilelist(path string) string {
	var strRet string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		strRet += path + "\r\n"
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return strRet
}

func main() {
	//flag.Parse()
	root := "F:\\Google\\Demo\\" //flag.Arg(0)
	println(getFilelist(root))
}
