package main

import (
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

func main() {
	router := httprouter.New()

	router.POST("/api/account/login", Login)
	router.POST("/api/account/signup", Signup)
	router.GET("/api/account/uname/:id", GetUserName)
	router.GET("/api/account/profile/:id", GetProfileImg)
	router.POST("/api/account/profile", RegisterProfile)
	router.POST("/api/img/upload", Upload)

	n := negroni.Classic()

	n.UseHandler(router)

	n.Run(":3000")
}

var (
	renderer     *render.Render
	mongoSession *mgo.Session
)

func init() {
	renderer = render.New()

	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	mongoSession = s
}
