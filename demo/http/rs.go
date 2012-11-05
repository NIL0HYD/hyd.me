package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/best", handler)
	//http.Handle("*", http.FileServer(http.Dir(".")))

	// we use StripPrefix so that /tmpfiles/somefile will access /img/somefile
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/img"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
