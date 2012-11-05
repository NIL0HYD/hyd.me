package main

import (
	"code.google.com/p/gorilla/mux"
	"hyd.me/webapp/controller"
	"log"
	"net/http"
)

// 所有的路由与相应的处理函数
var muxHandlers = map[string]func(rw http.ResponseWriter, req *http.Request){
	"/":             controller.Index,
	"/login":        controller.Login,
	"/logout":       controller.Logout,
	"/upload":       controller.Upload,
	"/findAllFiles": controller.FindAllFiles,
	"/mdeditor":     controller.Mdeditor,
	"/code":         controller.CodePrettify,
}

func main() {
	// http.Dir(".")代表当前路径
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	r := mux.NewRouter()
	for url, handler := range muxHandlers {
		r.HandleFunc(url, handler)
	}
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Print("Start the server.")
}
