package handle

import (
	"fmt"
	"mulbitchorong-backend/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
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
	ures := user.NewRes()
	uLogin := user.NewLogin()
	errs := binding.Bind(req, uLogin)
	if errs != nil {
		fmt.Println(errs)
	}

	session := mongoDB.Session.Copy()
	defer session.Close()

	c := session.DB("test").C("users")

	u := user.New()
	err := c.Find(bson.M{"uid": uLogin.ID}).One(&u)
	if err != nil {
		ures.Name = ""
		ures.IsSuccess = false
	} else {
		if u.UID == uLogin.ID {
			pwOK, _ := ComparePw(u.Pw, uLogin.Pw)
			if pwOK {
				ures.Name = u.Name
				ures.IsSuccess = true
			} else {
				ures.Name = ""
				ures.IsSuccess = false
			}
		} else {
			ures.Name = ""
			ures.IsSuccess = false
		}
	}

	renderer.JSON(w, http.StatusOK, ures)
}
