package handle

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// GetImage 는 filename에 해당하는 이미지를 응답합니다.
func GetImage(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	name := ps.ByName("filename")
	var img *ImageFile
	err := mongoDB.C("uploadfile", "fs.files").FindOne(context.TODO(), bson.M{"filename": name}).Decode(&img)
	if err != nil {
		log.Printf("GetImage : Failed to find %s", name)
		http.Error(w, "Failed to find "+name, http.StatusInternalServerError)
		return
	}
	bucket, _ := gridfs.NewBucket(
		mongoDB.DB("uploadfile"),
	)
	var buf bytes.Buffer
	_, err = bucket.DownloadToStreamByName(
		img.Filename,
		&buf,
	)
	if err != nil {
		log.Println("GetImage : Failed to open DownloadStream")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	if err != nil {
		log.Printf("Failed to read downloadStream %s", err)
		http.Error(w, "Failed to read downloadStream", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, req, name, time.Now(), bytes.NewReader(buf.Bytes()))
}
