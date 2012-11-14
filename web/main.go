package main

import (
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"html/template"
	"log"
	"net/http"
)

// 会话信息管理
var (
	globalSessions *session.Manager
)

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	log.Print("indexHandler...")
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	sess := globalSessions.SessionStart(rw, req)
	tmpl.Execute(rw, sess.Get("email"))
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	log.Printf("%s login", email)
	// password := req.FormValue("password")
	if req.Method == "GET" {
		return
	}
	sess := globalSessions.SessionStart(rw, req)
	sess.Set("email", email)
	// 跳转到首页
	http.Redirect(rw, req, "/", http.StatusFound)
}

// 注销登录
func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	log.Print("logout...")
	globalSessions.SessionDestroy(rw, req)
	// 跳转到首页
	http.Redirect(rw, req, "/", http.StatusFound)
}

func main() {
	// 记录日志
	/*
		wf, err := os.Create("log.md")
		if err != nil {
			return
		}
		log.SetOutput(wf)
	*/
	log.Print("Start the web server!")
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe is failed!")
	}
}
