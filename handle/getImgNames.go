package handle

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// GetImgNames 함수는 모든 이미지를 조회하여 모든 이미지의 파일이름을 응답하는 핸들러입니다.
func GetImgNames(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var imgfiles []*ImageFile
	cursor, err := mongoDB.C("uploadfile", "fs.files").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("GetImgNames: ", err)
	}
	if err := cursor.All(context.TODO(), &imgfiles); err != nil {
		log.Println("GetImgNames: ", err)
	}

	imgNames := make([]ImageName, len(imgfiles))
	for i := 0; i < len(imgfiles); i++ {
		imgNames[i].Name = imgfiles[i].Filename
	}
	renderer.JSON(w, http.StatusOK, imgNames)
}
