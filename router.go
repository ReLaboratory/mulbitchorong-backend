package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"gopkg.in/mgo.v2/bson"
)

// Signup 함수는 회원가입 기능을 수행하는 핸들러입니다.
func Signup(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	u := new(User)
	errs := binding.Bind(req, u)
	if errs != nil {
		fmt.Println(errs)
	}

	session := mongoSession.Copy()
	defer session.Close()

	u.ID = bson.NewObjectId()

	c := session.DB("test").C("users")

	if err := c.Insert(u); err != nil {
		renderer.JSON(w, http.StatusInternalServerError, err)
		return
	}

	ures := new(UserRes)
	ures.Name = u.Name
	if ures.Name != "" {
		ures.IsSuccess = true
	} else {
		ures.IsSuccess = false
	}

	renderer.JSON(w, http.StatusCreated, ures)
}

// Login 함수는 로그인 기능을 수행하는 핸들러입니다.
func Login(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {}

// GetUserName 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 이름을 응답하는 핸들러입니다.
func GetUserName(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB("test").C("users")

	result := User{}
	err := c.Find(bson.M{"uid": "testId"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User Name:", result.Name)
}

// GetProfileImg 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 프로필 이미지를 응답하는 핸들러입니다.
func GetProfileImg(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {}

// RegisterProfile 함수는 유저의 프로필 정보를 받아 프로필 데이터를 등록하는 핸들러입니다.
func RegisterProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {}

// Upload 함수는 유저 정보와 사용자 정보를 받아 이미지 업로드 기능을 수행하는 핸들러입니다.
func Upload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {}
