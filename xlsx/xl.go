package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	// xlsx-->sheets-->rows-->cells
	xl, err := xlsx.OpenFile("testfile.xlsx")
	if err != nil {
		return
	}

	sheetByName(xl)
	//allSheet(xl)
}

// 根据sheet名工作表
func sheetByName(xl *xlsx.File) {
	sheet := xl.Sheet["Tabelle1"]
	fmt.Println("Max: ", sheet.MaxRow, sheet.MaxCol)
	for i, row := range sheet.Rows {
		fmt.Printf("%v ", i)
		for _, cell := range row.Cells {
			fmt.Printf("%v %v ", cell.String(), cell.Value)
		}
		fmt.Printf("\n")
	}
}

// 所有的工作表
// 如果某个工作表中没有内容，遍历Cells则会panic
func allSheet(xl *xlsx.File) {
	for skey, sval := range xl.Sheet {
		fmt.Println(skey, sval)
		for rkey, rval := range sval.Rows {
			fmt.Printf("%v ", rkey)
			for _, cval := range rval.Cells {
				fmt.Printf("%v %v ", cval.String(), cval.Value)
			}
			fmt.Printf("\n")
		}
	}
}
