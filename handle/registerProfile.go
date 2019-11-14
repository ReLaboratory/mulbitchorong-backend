package handle

import (
	"fmt"
	"mulbitchorong-backend/user"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// RegisterProfile 함수는 유저의 프로필 정보를 받아 프로필 데이터를 등록하는 핸들러입니다.
func RegisterProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	registerRes := new(Res)

	req.ParseForm()
	_, fh, err := req.FormFile("files")
	uid := req.FormValue("uid")

	profileName := "profile_" + uid + "_mulbitchorong" + filepath.Ext(fh.Filename)

	session := mongoDB.Session.Copy()
	defer session.Close()

	c := session.DB("test").C("users")
	u := user.New()
	err = c.Find(bson.M{"uid": uid}).One(&u)
	if err != nil {
		fmt.Println("C.FIND ERROR: ", err)
	}
	err = c.Update(bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"profile_img": profileName}})
	if err != nil {
		fmt.Println("C.UPDATE ERROR: ", err)
	}
	file, _ := fh.Open()
	gridFile, err := session.DB("test").GridFS("fs").Create(profileName)
	if err != nil {
		registerRes.IsSuccess = false
	}
	fe := filepath.Ext(fh.Filename)
	fileExt := fe[1:]
	gridFile.SetMeta(bson.M{"uid": uid, "ext": fileExt})
	gridFile.SetName(profileName)

	if err := WriteToGridFile(file, gridFile); err != nil {
		registerRes.IsSuccess = false
	} else {
		registerRes.IsSuccess = true
	}
	renderer.JSON(w, http.StatusCreated, registerRes)
}
