package service

import (
	"errors"
	"log"
)

// 1. 定义可使用的服务
// 直接开放给client使用
type Service interface {
	Set(key, value interface{})      // 设置map中内容
	Get(key interface{}) interface{} // 获取map中内容
}

// 2. 定义服务提供者可提供的服务
type Provider interface {
	GetService() Service
}

// 3. 注册服务提供者
var provides = make(map[string]Provider)

// 注册方法
func Register(name string, provider Provider) {
	log.Print("Register provider ", name)
	if provider == nil {
		panic("Provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("Register called twice for provider " + name)
	}
	provides[name] = provider
}

// 4. 访问服务提供者的服务
type Accesser struct {
	provider Provider
}

func NewAccesser(name string) (*Accesser, error) {
	log.Print("New Accesser.")
	provider, ok := provides[name]
	if !ok {
		return nil, errors.New("Unknow provider")
	}
	return &Accesser{provider}, nil
}

// 获取服务
func (a *Accesser) GetService() Service {
	return a.provider.GetService()
}
