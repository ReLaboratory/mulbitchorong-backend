package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// 라우터 생성
	router := httprouter.New()

	// 핸들러 정의
	router.GET("/api/account/uname/:id", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		fmt.Printf("현재 유저의 Id 값은 %s입니다.\n", ps.ByName("id"))
	})

	// negroni 미들웨어 생성
	n := negroni.Classic()

	// negroni에 router를 핸들러로 등록
	n.UseHandler(router)

	// 서버 실행
	n.Run(":3000")
}
