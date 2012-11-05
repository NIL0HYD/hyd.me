package main

import (
	"bufio"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("F:\\Google\\Demo\\ch.go")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		//n, _ := f.Read(buf)
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[0:n])
		/*os.Stdout.Write(buf[:n])*/
	}
}
