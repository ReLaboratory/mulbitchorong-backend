package handle

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Signup 함수는 회원가입 기능을 수행하는 핸들러입니다.
func Signup(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	statusCode := http.StatusCreated

	u := new(User)
	errs := binding.Bind(req, u)
	if errs != nil {
		fmt.Println(errs)
	}
	u.ID = bson.NewObjectId()

	ures := new(UserRes)
	ures.Name = u.Name

	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB("test").C("users")

	IDCheck := User{}
	err := c.Find(bson.M{"uid": u.UID}).One(&IDCheck)
	if err != nil {
		ures.IsSuccess = true

		hashedPw, _ := bcrypt.GenerateFromPassword([]byte(u.Pw), bcrypt.DefaultCost)
		u.Pw = string(hashedPw[:])

		if err := c.Insert(u); err != nil {
			renderer.JSON(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		ures.IsSuccess = false
		ures.Name = ""
	}

	if !ures.IsSuccess {
		statusCode = http.StatusOK
	}

	renderer.JSON(w, statusCode, ures)
}
