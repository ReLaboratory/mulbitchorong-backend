package handle

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// GetImageUser 는 해당 파일을 업로드한 user의 id를 응답합니다.
func GetImageUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var imgfile *ImageFile
	filename := ps.ByName("filename")
	err := mongoDB.C("uploadfile", "fs.files").FindOne(context.TODO(), bson.M{"filename": filename}).Decode(&imgfile)
	if err != nil {
		fmt.Println("GetImageUser : ", err)
	}
	imgUploader := new(ImageUploader)
	imgUploader.UID = imgfile.MetaData.UID
	renderer.JSON(w, http.StatusOK, imgUploader)
}
