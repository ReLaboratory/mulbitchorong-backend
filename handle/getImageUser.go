package handle

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetImageUser 는 해당 파일을 업로드한 user의 id를 응답합니다.
func GetImageUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	filename := ps.ByName("filename")

	session := mongoDB.Session.Copy()
	defer session.Close()

	c := session.DB("test").C("fs" + ".files")

	var imgfile *ImageFile
	err := c.Find(filename).One(&imgfile)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	imgUploader := new(ImageUploader)

	imgUploader.UID = imgfile.MetaData.UID
	renderer.JSON(w, http.StatusOK, imgUploader)
}
