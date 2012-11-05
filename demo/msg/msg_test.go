package msg

import (
	"testing"
)

func TestMsg(t *testing.T) {
	m := Msg{"Hello", "World"}
	m.Sent()
}
