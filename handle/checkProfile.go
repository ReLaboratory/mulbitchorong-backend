package handle

import (
	"fmt"
	"mulbitchorong-backend/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// CheckProfile 함수는 프로필 등록여부를 확인하여 응답하는 핸들러입니다.
func CheckProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uid := ps.ByName("id")
	res := new(Res)

	session := mongoDB.Session.Copy()
	defer session.Close()

	c := session.DB("test").C("users")
	u := user.New()
	err := c.Find(bson.M{"uid": uid}).One(&u)

	if err != nil {
		res.IsSuccess = false
	}

	name := u.ProfileImg

	if name != "" {
		fmt.Println("NAME: ", name)
		res.IsSuccess = true
	} else {
		fmt.Println("NAME: ", name)
		res.IsSuccess = false
	}
	renderer.JSON(w, http.StatusOK, res)
}
