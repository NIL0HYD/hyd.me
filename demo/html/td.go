package main

import (
	"log"
	"os"
	"text/template"
)

const tt = `
{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}
`

func exec(tmpl *template.Template) {
	tmpl.Execute(os.Stdout, "data")
}

func execTemp(tmpl *template.Template) {
	tmpl.ExecuteTemplate(os.Stdout, "T3", "data")
}

func main() {
	tmpl, err := template.New("test").Parse(tt)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	exec(tmpl)

	execTemp(tmpl)
}
