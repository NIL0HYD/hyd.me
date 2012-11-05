package main

import "fmt"
import "container/list"

// 这是不是list.Remove的bug
// https://groups.google.com/forum/?fromgroups=#!searchin/golang-china/list$20remove/golang-china/vhXBQHizVC8/JH1wM49XRwEJ
func PrintList(l *list.List) {
	fmt.Println("\tLen:", l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("\t\t[%d] %d\n", i, e.Value)
		i++
	}
}

func main() {
	l1, l2 := list.New(), list.New()

	l1_e1, l1_e2, l1_e3 := l1.PushBack(1), l1.PushBack(2), l1.PushBack(3)
	l2_e1, l2_e2, l2_e3 := l2.PushBack(1), l2.PushBack(2), l2.PushBack(3)

	fmt.Println("List 1")
	PrintList(l1)
	fmt.Printf("%d, %d, %d\n", l1_e1.Value, l1_e2.Value, l1_e3.Value)
	fmt.Println("List 2")
	PrintList(l2)
	fmt.Printf("%d, %d, %d\n", l2_e1.Value, l2_e2.Value, l2_e3.Value)

	// 从l1中删除l2_e2
	fmt.Println("Delete l2_e2 from l1")
	l1.Remove(l2_e2)
	fmt.Println("List 1")
	PrintList(l1)
	fmt.Println("List 2")
	PrintList(l2)

	// 在l2的l2_e2前插入一个值8，这个不会成功
	fmt.Println("Insert value 8 before l2_e2")
	l2.InsertBefore(8, l2_e2)
	fmt.Println("List 1")
	PrintList(l1)
	fmt.Println("List 2")
	PrintList(l2)
}
