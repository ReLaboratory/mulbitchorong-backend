package handle

import (
	"context"
	"log"
	"mulbitchorong-backend/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUserName 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 이름을 응답하는 핸들러입니다.
func GetUserName(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	result := user.New()
	err := mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": ps.ByName("id")}).Decode(&result)

	uname := user.NewName()
	if err != nil {
		uname.Name = ""
		log.Println("GetUserName : Failed to find user ", ps.ByName("id"))
	} else {
		uname.Name = result.Name
		log.Printf("GetUserName : User %s Find Success ", uname.Name)
	}

	renderer.JSON(w, http.StatusOK, uname)
}
