package main

import (
	"html/template"
	"log"
	"net/http"
)

type Myinfo struct {
	Name  string
	Email string
	QQ    string
}

func handler(w http.ResponseWriter, r *http.Request) {

	myinfo := Myinfo{"黄远董", "yuandong.huang@gmail.com", "176458203"}
	//Title := "Hello, 世界！"
	tpl, _ := template.ParseFiles("edit.html")
	err := tpl.Execute(w, myinfo)
	if err != nil {
		log.Fatalf("exec: %s", err)
	}
}

// Prepare some data to insert into the template.
type Recipient struct {
	Name, Gift string
	Attended   bool
}

var recipients = []Recipient{
	{"Aunt Mildred", "bone china tea set", true},
	{"Uncle John", "moleskin pants", false},
	{"Cousin Rodney", "", false},
}

func sliceHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("sl.html")
	for _, r := range recipients {
		err := tpl.Execute(w, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

func sllHandler(w http.ResponseWriter, r *http.Request) {
	sl := []string{"QQQ", "AAA", "BBB"}
	tpl, _ := template.ParseFiles("sll.html")
	err := tpl.Execute(w, sl)
	if err != nil {
		log.Println("executing template:", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sl", sliceHandler)
	http.HandleFunc("/sll", sllHandler)
	http.ListenAndServe(":8080", nil)
}
