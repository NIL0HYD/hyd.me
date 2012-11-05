package controller

import (
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"html/template"
	"hyd.me/webapp/service"
	"log"
	"net/http"
)

// 全局变量
var globalSessions *session.Manager

// 初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

// 首页
func Index(rw http.ResponseWriter, req *http.Request) {

	log.Print(req.RemoteAddr)
	log.Print("Here is index...")

	tpl := template.Must(template.ParseFiles("html/index.html", "html/header.html", "html/footer.html"))
	// session管理
	session := globalSessions.SessionStart(rw, req)
	tpl.ExecuteTemplate(rw, "content", session.Get("profile"))
	//tpl.Execute(rw, session.Get("profile"))
}

// 登录页
func Login(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		log.Print("Here GET")
		tpl := template.Must(template.ParseFiles("html/login.html", "html/header.html", "html/footer.html"))
		tpl.ExecuteTemplate(rw, "content", nil)
	} else {

		log.Print("Here POST")

		email := req.FormValue("email")
		password := req.FormValue("password")

		log.Print(password)

		// 根据用户名获取用户基本信息
		result := service.GetUser(email)
		session := globalSessions.SessionStart(rw, req)
		session.Set("profile", result)

		log.Print(session.Get("profile"))

		http.Redirect(rw, req, "/", http.StatusFound)
	}
}

// 退出
func Logout(rw http.ResponseWriter, req *http.Request) {
	log.Print("Logout...")
	globalSessions.SessionDestroy(rw, req)

	http.Redirect(rw, req, "/", http.StatusFound)
}

// Markdown editer
func Mdeditor(rw http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("html/mdeditor.html"))
	tpl.Execute(rw, nil)
}

// google-code-prettify
func CodePrettify(rw http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("html/code.html"))
	tpl.Execute(rw, nil)
}
