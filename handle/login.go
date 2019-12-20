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

// ComparePw 함수는 해쉬화된 Pw와 평문 Pw를 비교하는 함수입니다.
func ComparePw(hash, pw string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Login 함수는 로그인 기능을 수행하는 핸들러입니다.
func Login(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uLogin := user.NewLogin()
	err := binding.Bind(req, uLogin)
	if err != nil {
		fmt.Println(err)
	}

	u := user.New()
	ures := user.NewRes()
	err = mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": uLogin.ID}).Decode(&u)

	if err != nil {
		ures.Name = ""
		ures.IsSuccess = false
		log.Println("Login : Login failed ", uLogin.ID, " != ", u.UID)
		log.Println("Login : ", err)
	} else {
		if u.UID == uLogin.ID {
			pwOK, _ := ComparePw(u.Pw, uLogin.Pw)
			if pwOK {
				ures.Name = u.Name
				ures.IsSuccess = true
				log.Printf("Login : User %s Login Successful", ures.Name)
			} else {
				ures.Name = ""
				ures.IsSuccess = false
				log.Println("Login : Login failed")
			}
		} else {
			ures.Name = ""
			ures.IsSuccess = false
			log.Println("Login : Login failed")
		}
	}

	renderer.JSON(w, http.StatusOK, ures)
}
