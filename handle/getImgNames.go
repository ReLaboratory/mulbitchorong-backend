package handle

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetImgNames 함수는 모든 이미지를 조회하여 모든 이미지의 파일이름을 응답하는 핸들러입니다.
func GetImgNames(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	session := mongoDB.Session.Copy()
	defer session.Close()

	c := session.DB("test").C("fs" + ".files")

	var imgfiles []*ImageFile
	err := c.Find(nil).All(&imgfiles)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	imgNames := make([]ImageName, len(imgfiles))
	for i := 0; i < len(imgfiles); i++ {
		imgNames[i].Name = imgfiles[i].Filename
	}
	renderer.JSON(w, http.StatusOK, imgNames)
}
