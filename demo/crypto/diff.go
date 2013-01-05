package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// md5
	h := md5.New()
	h.Write([]byte("404 Pages"))
	//fmt.Printf("%x \n", h.Sum(nil))
	s := "0B2441996dc4ebeb1a356ebfe1b3e9a5"

	fmt.Println(fmt.Sprintf("%x", h.Sum(nil)))

	if strings.EqualFold(s, fmt.Sprintf("%x", h.Sum(nil))) {
		fmt.Println("They are equals! EqualFold")
	}

	if s == fmt.Sprintf("%x", h.Sum(nil)) {
		fmt.Println("They are equals!")
	}

	// sha-1
	h = sha1.New()
	h.Write([]byte("404 Pages"))
	fmt.Printf("%x \n", h.Sum(nil))

	// 文件hash
	f := md5.New()
	file := "F:\\cmdlet.txt"
	r, err := os.Open(file)
	if err != nil {
		return
	}
	defer r.Close()
	b, _ := ioutil.ReadAll(r)
	f.Write(b)
	fmt.Printf("File MD5: %x \n", f.Sum(nil))
}
