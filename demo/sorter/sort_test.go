package sorter

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	ints := IntSorter{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := StringSorter{"nut", "ape", "elephant", "zoo", "go"}
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
	//t.Fail()
}
