package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//read()
	//write()
	//quickRead()
	//readFor()
	//writeString()
	//copy()
	//copyN()
	bufiowr()
}

func bufiowr() {
	buf := make([]byte, 1024)
	f, err := os.Open("copy.go")
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	wf, err := os.Create("copy.md")
	w := bufio.NewWriter(wf)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[0:n])
	}
}

func quickRead() {
	bs, err := ioutil.ReadFile("file.go")
	if err != nil {
		return
	}
	fmt.Println(string(bs))
}

func copyN() {
	rf, err := os.Open("dir.go")
	if err != nil {
		return
	}
	defer rf.Close()

	if stat, err := rf.Stat(); err == nil {
		fmt.Println(stat.Name(), stat.Size())
	}

	wf, err := os.Create("dirnew.md")
	if err != nil {
		return
	}
	defer wf.Close()

	_, err = io.CopyN(wf, rf, 256)
	if err != nil {
		return
	}
}

func copy() {
	rf, err := os.Open("dir.go")
	if err != nil {
		return
	}
	defer rf.Close()

	wf, err := os.Create("dirnew.md")
	if err != nil {
		return
	}
	defer wf.Close()

	_, err = io.Copy(wf, rf)
	if err != nil {
		return
	}
}

func readFor() {
	rf, err := os.Open("dir.go")
	if err != nil {
		return
	}
	defer rf.Close()
	bs := make([]byte, 512)
	for {
		n, err := rf.Read(bs)
		if err != nil {
			return
		}
		if n == 0 {
			break
		}
		fmt.Println(string(bs[0:n]))
	}
}

func read() {
	rf, err := os.Open("dir.go")
	if err != nil {
		return
	}
	defer rf.Close()
	stat, err := rf.Stat()
	if err != err {
		return
	}
	bs := make([]byte, stat.Size())
	_, err = rf.Read(bs)
	fmt.Println(string(bs))
}

func write() {
	bs := []byte("# Hello, World!")
	wf, err := os.Create("ionew.md")
	if err != nil {
		return
	}
	defer wf.Close()

	wf.Write(bs)
}

func writeString() {

	wf, err := os.Create("since.md")
	if err != nil {
		return
	}
	defer wf.Close()

	wf.WriteString("# since 1986")
}
