package user

import (
	"net/http"

	"github.com/mholt/binding"
	"gopkg.in/mgo.v2/bson"
)

// User 구조체는 유저의 전체 정보를 저장합니다.
type User struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	UID  string        `bson:"uid" json:"uid"`
	Pw   string        `bson:"pw" json:"pw"`
	Name string        `bson:"uname" json:"uname"`
}

// Res 구조체는 생성된 유저의 이름과 회원가입 성공 여부를 저장합니다.
type Res struct {
	Name      string `json:"uname"`
	IsSuccess bool   `json:"isSuccess"`
}

// Login 구조체는 로그인을 수행하는 유저의 정보를 담고 있습니다.
type Login struct {
	ID string `json:"uid"`
	Pw string `json:"pw"`
}

// Name 구조체는 유저의 이름을 담고 있습니다.
type Name struct {
	Name string `json:"uname"`
}

// FieldMap 메서드는 Login 타입을 binding.FieldMapper 인터페이스이도록 하기 위해 만든 메서드입니다.
func (u *Login) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&u.ID: "uid",
		&u.Pw: "pw",
	}
}

// FieldMap 메서드는 User 타입을 binding.FieldMapper 인터페이스이도록 하기 위해 만든 메서드입니다.
func (u *User) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&u.UID:  "uid",
		&u.Pw:   "pw",
		&u.Name: "uname",
	}
}

// New 함수는 새로운 User 구조체를 생성합니다.
func New() *User {
	u := new(User)
	return u
}

// NewRes 함수는 새로운 Res 구조체를 생성합니다.
func NewRes() *Res {
	ures := new(Res)
	return ures
}

// NewLogin 함수는 새로운 User 구조체를 생성합니다.
func NewLogin() *Login {
	u := new(Login)
	return u
}

// NewName 함수는 새로운 Name 구조체를 생성합니다.
func NewName() *Name {
	u := new(Name)
	return u
}
