package provider

import (
	"hyd.me/service"
	"log"
)

//
type Service struct {
	Profile map[interface{}]interface{}
}

func (s *Service) Get(key interface{}) interface{} {
	return s.Profile[key]
}

func (s *Service) Set(key, value interface{}) {
	s.Profile[key] = value
}

//
type Provider struct {
	service *Service
}

func (p *Provider) GetService() service.Service {
	return p.service
}

var provider = &Provider{&Service{make(map[interface{}]interface{})}}

func init() {
	log.Print("Init & Register Provider")
	service.Register("nil0hyd", provider)
}
