package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// 配置信息
var config map[string]string

// 初始化,读取配置文件
func init() {
	file, err := os.Open("hyd.conf")
	if err != nil {
		log.Print("read config file failed!")
		panic(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	err = dec.Decode(&config)

	if err != nil {
		log.Print("read config file failed!")
		panic(err)
	}
}

func main() {
	fmt.Println(config["name"])
	fmt.Println(config["qq"])

}
