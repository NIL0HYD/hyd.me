package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tmpls = template.Must(template.ParseFiles("index.html",
		"jsmethods/js_oo.html", "convertRMB/convertRMB.html"))
)

func index(rw http.ResponseWriter, req *http.Request) {
	renderHtml(rw, "index.html")
}

func jsmethods(rw http.ResponseWriter, req *http.Request) {
	renderHtml(rw, "js_oo.html")
}

func convertRMB(rw http.ResponseWriter, req *http.Request) {
	renderHtml(rw, "convertRMB.html")
}

func renderHtml(rw http.ResponseWriter, tmpl string) {
	err := tmpls.ExecuteTemplate(rw, tmpl, nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/rmb", convertRMB)
	http.HandleFunc("/jsmethods", jsmethods)

	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/jsmethods/", http.FileServer(http.Dir(".")))
	http.Handle("/convertRMB/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Print("Start the server.")
}
