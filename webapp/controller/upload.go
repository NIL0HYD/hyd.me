package controller

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"hyd.me/webapp/service"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 文件上传
func Upload(rw http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		log.Print("upload GET...")
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("html/upload.html")
		t.Execute(rw, token)
		return
	}

	log.Print("upload file...")
	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(rw, "%v", handler.Header)

	//saveFileSystem(file, handler.Filename)

	saveMongodb(file, handler.Filename)

}

// 保存到文件系统
func saveFileSystem(file multipart.File, fileName string) {
	f, err := os.OpenFile("upload/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return
	}
}

// 保存到Mongodb
func saveMongodb(file multipart.File, fileName string) {
	log.Print("Saving the file to Mongodb!")
	err := service.CreateFile(file, fileName)
	if err != nil {
		return
	}
}

// 查找所有的文件
func FindAllFiles(rw http.ResponseWriter, req *http.Request) {
	service.FindAllFiles()
}
