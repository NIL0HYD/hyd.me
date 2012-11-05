package main

import (
	"bytes"
	"encoding/json"
	"os"
)

type Message struct {
	Name, Text string
}

func main() {
	src, _ := json.Marshal(&Message{"annnnn", "extt<tt>"})
	var b bytes.Buffer
	//json.Compact(&b, src)
	json.HTMLEscape(&b, src)
	b.WriteTo(os.Stdout)
}
