package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 定义一个访问者结构体
type Visitor struct{}

var contents string

var dir = "E:\\Finance-OA\\workspace\\CymcoMobileOA.v2.0\\src\\com\\renda\\soft\\service\\impl\\"

func writeString(str string) {
	contents += str
	//contents := fmt.Fprintf(w, format, ...)
}

func (self *Visitor) visit(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	} else if (f.Mode() & os.ModeSymlink) > 0 {
		return nil
	} else {
		if strings.HasSuffix(f.Name(), ".java") {
			fmt.Println(f)
			writeString("源代码文件：" + f.Name() + "\n")
			file, err := os.Open(dir + f.Name())
			if err != nil {
				return err
			}
			defer file.Close()
			input, err2 := ioutil.ReadAll(file)

			if err2 != nil {
				fmt.Printf("%v\n", err2)
			}
			// 合并
			writeString(string(input))
		}
	}
	return nil
}

func main() {
	v := &Visitor{}

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		return v.visit(path, f, err)
	})

	// 写入的文件
	wf, err := os.Create("F:\\com.renda.soft.service.impl.doc")
	if err != nil {
		return
	}
	defer wf.Close()

	wf.WriteString(contents)

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
