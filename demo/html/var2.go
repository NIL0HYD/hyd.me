package main

import (
	"os"
	"reflect"
	"text/template"
)

var fns = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

func main() {
	text := `{{range  $i, $e := .}} {{$e}} {{if last $i $e | not}}, {{end}}{{end}}.`
	t := template.Must(template.New("abc").Funcs(fns).Parse(text))
	a := []string{"Alice", "Bob", "Cherry", "Dark"}
	t.Execute(os.Stdout, a)
}
