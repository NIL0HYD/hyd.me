package service

import (
	"fmt"
	"io"
	"labix.org/v2/mgo"
	"log"
	"mime/multipart"
	"os"
)

// 保存文件到DB
func CreateFile(reader multipart.File, fileName string) error {
	log.Print("Service CreateFile ...")
	writer, err := db.GridFS("fs").Create(fileName)
	if err != nil {
		return fmt.Errorf("Mongo create file: %v error!", fileName)
	}
	defer writer.Close()
	_, err = io.Copy(writer, reader)
	if err != nil {
		return fmt.Errorf("Copy file: %v error!", fileName)
	}
	return nil
}

// 查找文件
func FindAllFiles() {
	gfs := db.GridFS("fs")
	iter := gfs.Find(nil).Iter()
	var f *mgo.GridFile
	for gfs.OpenNext(iter, &f) {
		fmt.Printf("Filename: %s, %v\n", f.Name(), f.Id())
		b := make([]byte, 512)
		w, err := os.Create("upload/" + f.Name())
		if err != nil {
			continue
		}
		for {
			n, _ := f.Read(b)
			if n == 0 {
				break
			}
			w.Write(b[0:n])
		}
		defer w.Close()
	}
	if iter.Err() != nil {
		panic(iter.Err())
	}
}
