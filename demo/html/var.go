package main

import (
	"os"
	"text/template"
)

func main() {
	text := `{{ range  $i, $e := . }}{{ if $i }}, {{ end }}{{ $e }}{{ end }}.`
	t := template.Must(template.New("abc").Parse(text))
	a := []string{"Alice", "Bob", "Cherry", "Dark"}
	t.Execute(os.Stdout, a)
}
