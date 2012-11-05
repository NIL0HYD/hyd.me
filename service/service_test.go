package service

import (
	"hyd.me/service"
	_ "hyd.me/service/provider"
	"log"
	"testing"
)

var accesser *service.Accesser

func TestService(t *testing.T) {
	log.Print("Here test running...")
	accesser, err := service.NewAccesser("nil0hyd")
	if err != nil {
		return
	}
	serv := accesser.GetService()
	serv.Set("name", "HuangYuanDong")

	if serv.Get("name") != "HuangYuanDong" {
		t.Fail()
	}
	log.Print(serv.Get("name"))
}
