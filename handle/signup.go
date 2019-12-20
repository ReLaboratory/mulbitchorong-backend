package handle

import (
	"context"
	"fmt"
	"log"
	"mulbitchorong-backend/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Signup 함수는 회원가입 기능을 수행하는 핸들러입니다.
func Signup(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	u := user.New()
	errs := binding.Bind(req, u)
	if errs != nil {
		fmt.Println(errs)
	}

	ures := user.NewRes()
	ures.Name = u.Name
	IDCheck := user.User{}
	err := mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": u.UID}).Decode(&IDCheck)
	if err != nil {
		ures.IsSuccess = true

		hashedPw, _ := bcrypt.GenerateFromPassword([]byte(u.Pw), bcrypt.DefaultCost)
		u.Pw = string(hashedPw[:])

		if _, err := mongoDB.C("test", "users").InsertOne(context.TODO(), u); err != nil {
			log.Println("Signup : ", err)
			renderer.JSON(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		ures.IsSuccess = false
		ures.Name = ""
		log.Printf("Signup : User %s is already registered", IDCheck.UID)
	}

	statusCode := http.StatusCreated
	if !ures.IsSuccess {
		statusCode = http.StatusOK
	}

	renderer.JSON(w, statusCode, ures)
}
