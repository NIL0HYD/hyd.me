package myinfo

import (
	"fmt"
	"testing"
)

func TestGetEmail(t *testing.T) {
	fmt.Println(GetEmail())
	t.Fail()
}
