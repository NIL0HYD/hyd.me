package main

import (
	"os"
	"text/template"
)

var fns = template.FuncMap{
	"eq": func(x, y interface{}) bool {
		return x == y
	},
	"sub": func(y, x int) int {
		return x - y
	},
}

func main() {
	// In a chained pipeline, 
	// the result of the each command is passed as the last argument of the following command.
	// 在一个管道链中，每个命令的结果被传递到下一个命令做为最后一个参数
	text := `{{range  $i, $e := .}} {{$e}} {{if len $ | sub 1 | eq $i | not}}, {{end}}{{end}}.`
	t := template.Must(template.New("abc").Funcs(fns).Parse(text))
	a := []string{"Alice", "Bob", "Cherry", "Dark"}
	t.Execute(os.Stdout, a)
}
