package msg

import (
	"fmt"
)

type Msg struct {
	Content  string
	Receiver string
}

func (m *Msg) Sent() {
	fmt.Printf("%s, %s!\n", m.Content, m.Receiver)
}
