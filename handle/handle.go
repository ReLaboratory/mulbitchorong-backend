package handle

import (
	"log"
	"net/http"
	"os/user"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// GetUserName 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 이름을 응답하는 핸들러입니다.
func GetUserName(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB("test").C("users")

	result := user.User{}
	err := c.Find(bson.M{"uid": ps.ByName("id")}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	uname := UserName{}
	uname.Name = result.Name
	renderer.JSON(w, http.StatusOK, uname)
}

// GetProfileImg 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 프로필 이미지를 응답하는 핸들러입니다.
func GetProfileImg(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {}

// RegisterProfile 함수는 유저의 프로필 정보를 받아 프로필 데이터를 등록하는 핸들러입니다.
func RegisterProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {}

// Upload 함수는 유저 정보와 사용자 정보를 받아 이미지 업로드 기능을 수행하는 핸들러입니다.
func Upload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {}
