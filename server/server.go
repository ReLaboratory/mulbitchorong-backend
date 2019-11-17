package server

import (
	"mulbitchorong-backend/handle"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

// Server 객체는 Server에 대한 정보를 담고 있습니다.
type Server struct {
	neg    *negroni.Negroni
	router *httprouter.Router
}

// New 함수는 Server에 대한 설정을 담당합니다
func New() (*Server, error) {
	sv := new(Server)
	sv.router = httprouter.New()
	sv.neg = negroni.Classic()

	err := handle.InitMongo("mongodb://localhost")
	if err != nil {
		return nil, err
	}

	sv.router.POST("/api/account/login", handle.Login)
	sv.router.POST("/api/account/signup", handle.Signup)
	sv.router.GET("/api/account/uname/:id", handle.GetUserName)
	sv.router.GET("/api/account/profile/:id", handle.GetProfileImg)
	sv.router.POST("/api/account/profile", handle.RegisterProfile)
	sv.router.PUT("/api/account/profile", handle.RegisterProfile)
	sv.router.GET("/api/account/profile-registered/:id", handle.CheckProfile)
	sv.router.POST("/api/img/upload", handle.Upload)
	sv.router.GET("/api/img/upload-name", handle.GetImgNames)
	sv.router.GET("/api/img/upload-file/:filename", handle.GetImage)
	sv.router.GET("/api/img/upload-user/:filename", handle.GetImageUser)

	sv.neg.UseHandler(sv.router)
	return sv, nil
}

// Run 함수는 Server를 실행합니다.
func (s *Server) Run(port string) {
	s.neg.Run(port)
}
