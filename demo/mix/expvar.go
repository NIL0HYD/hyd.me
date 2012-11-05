package main

import (
	"expvar"
	"fmt"
)

func main() {
	fmt.Println("Hello, Go1. This is liigo.")
	expvar.Do(func(kv expvar.KeyValue) { fmt.Printf("\n%s=%s\n", kv.Key, kv.Value) })
}
