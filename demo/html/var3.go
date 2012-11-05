package main

import (
	"os"
	"strings"
	"text/template"
)

var fns = template.FuncMap{
	"join": strings.Join,
}

func main() {
	t := template.Must(template.New("abc").Funcs(fns).Parse(`{{join . ", "}}.`))
	a := []string{"Alice", "Bob", "Cherry", "Dark"}
	t.Execute(os.Stdout, a)
}
