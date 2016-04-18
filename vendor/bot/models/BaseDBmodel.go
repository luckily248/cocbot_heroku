package models

import (
	"gopkg.in/mgo.v2"
)

type BaseDBmodel struct {
	session *mgo.Session
	db      *mgo.Database
	c       *mgo.Collection
}

func (this *BaseDBmodel) DBname() string {
	return "dataserver"
}

//成功初始化后必须调用  defer this.session.Close()
func (this *BaseDBmodel) init() (err error) {
	newsession, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		return
	}
	this.session = newsession
	this.session.SetMode(mgo.Monotonic, true)
	this.db = this.session.DB(this.DBname())
	return
}

func (this *BaseDBmodel) Check() (err error) {
	err = this.init()
	if err != nil {
		return
	}
	defer this.session.Close()
	return
}
