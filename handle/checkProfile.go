package handle

import (
	"context"
	"log"
	"mulbitchorong-backend/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckProfile 함수는 프로필 등록여부를 확인하여 응답하는 핸들러입니다.
func CheckProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uid := ps.ByName("id")
	res := new(Res)
	u := user.New()
	err := mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": uid}).Decode(&u)

	if err != nil {
		res.IsSuccess = false
		log.Println("CheckProfile : IsSuccess false")
		log.Println("CheckProfile : ", err)
		renderer.JSON(w, http.StatusOK, res)
		return
	}

	name := u.ProfileImg

	if name != "" {
		res.IsSuccess = true
		log.Println("CheckProfile : IsSuccess true")
	} else {
		res.IsSuccess = false
		log.Println("CheckProfile : IsSuccess false")
	}
	renderer.JSON(w, http.StatusOK, res)
}
