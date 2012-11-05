package main

import (
	"crypto/aes"
	. "fmt"
	"os"
)

func main() {
	msg := "My name is Astaxie"
	// some key, 16 Byte long
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	Println("len of message: ", len(msg))
	Println("len of key: ", len(key))
	// create the new cipher
	// 参数key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法
	c, err := aes.NewCipher(key)
	if err != nil {
		Println("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}

	out := make([]byte, len(msg))

	c.Encrypt(out, []byte(msg)) // encrypt the first half
	//c.Encrypt(msgbuf[16:32], out[16:32]) // encrypt the second half

	Println("len of encrypted: ", len(out))
	Println(">> ", out)

	// // now we decrypt our encrypted text
	plain := make([]byte, len(out))
	c.Decrypt(plain, out)

	Println("msg: ", string(plain))
}
