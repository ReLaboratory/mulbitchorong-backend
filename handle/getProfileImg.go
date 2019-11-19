package handle

import (
	"log"
	"mulbitchorong-backend/user"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// GetProfileImg 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 프로필 이미지를 응답하는 핸들러입니다.
func GetProfileImg(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uid := ps.ByName("id")

	session := mongoDB.Session.Copy()
	defer session.Close()

	c := session.DB("test").C("users")
	u := user.New()
	err := c.Find(bson.M{"uid": uid}).One(&u)

	name := u.ProfileImg

	gridfs := mongoDB.Session.DB("test").GridFS("fs")

	f, err := gridfs.Open(name)
	if err != nil {
		if name == " " {
			log.Printf("Failed to open %s: %v", name, err)
		} else {
			log.Printf("ProfileImage is emty")
		}
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	http.ServeContent(w, req, name, time.Now(), f)
}
