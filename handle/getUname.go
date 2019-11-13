package handle

import (
	"log"
	"mulbitchorong-backend/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// GetUserName 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 이름을 응답하는 핸들러입니다.
func GetUserName(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	session := mongoDB.Session.Copy()
	defer session.Close()

	c := session.DB("test").C("users")

	result := user.New()
	err := c.Find(bson.M{"uid": ps.ByName("id")}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	uname := user.NewName()
	uname.Name = result.Name

	renderer.JSON(w, http.StatusOK, uname)
}
