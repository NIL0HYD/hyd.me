package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	salt2()
	//salt()
	//normal()
}

func normal() {
	h := md5.New()
	h.Write([]byte("404 Pages"))
	fmt.Printf("%x \n", h.Sum(nil))
	//9ee3 2fbc 5704 75f6 11d9 7078 c482 7fa4
	//6dcd 4d36 fd97 ba5f 119c 000d 6be9 3439
	//a87ff679a2f3e71d9181a67b7542122c

	h = sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x \n", h.Sum(nil))

	h = sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x \n", h.Sum(nil))

	h = md5.New()
	io.WriteString(h, "需要加密的密码")
	fmt.Printf("%x \n", h.Sum(nil))
}

func salt() {
	//import "crypto/md5"
	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "需要加密的密码")

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
}

func salt2() {
	//import "crypto/md5"
	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "1{sysadmin}")

	//pwmd5等于5536484b09fdfa322c70338849bcff11
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(pwmd5)
}
