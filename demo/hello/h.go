package main

import (
	"hyd.me/msg"
)

func main() {
	m := msg.Msg{"Nice to meet you", "hyd"}
	m.Sent()
}
