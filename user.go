package main

import (
	"net/http"

	"github.com/mholt/binding"
	"gopkg.in/mgo.v2/bson"
)

// User 구조체는 사용자 정보를 저장합니다.
type User struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	UID  string        `bson:"uid" json:"uid"`
	Pw   string        `bson:"pw" json:"pw"`
	Name string        `bson:"uname" json:"uname"`
}

// FieldMap 메서드는 User 타입을 binding.FieldMapper 인터페이스이도록 하기 위해 만든 메서드입니다.
func (u *User) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&u.UID:  "uid",
		&u.Pw:   "pw",
		&u.Name: "uname",
	}
}
