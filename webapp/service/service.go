package service

import (
	"hyd.me/webapp/entity"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// 常量
const (
	dbname = "hyd"
	dburl  = "root:nice@127.0.0.1:27017/hyd"
)

// MongoDB
var db *mgo.Database

func init() {
	// MongoDB 初始化
	session, err := mgo.Dial(dburl)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db = session.DB(dbname)
}

// 根据用户名获取用户基本信息
func GetUser(email string) *entity.Profile {
	// 查找个人信息
	c := db.C("profile")
	result := &entity.Profile{}
	err := c.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}
