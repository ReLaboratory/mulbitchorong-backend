package handle

import (
	"bytes"
	"context"
	"log"
	"mulbitchorong-backend/user"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// GetProfileImg 함수는 유저의 ID값에 맞는 데이터를 조회하여 해당하는 유저의 프로필 이미지를 응답하는 핸들러입니다.
func GetProfileImg(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uid := ps.ByName("id")

	u := user.New()
	err := mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": uid}).Decode(&u)

	name := u.ProfileImg
	var img *ImageFile
	err = mongoDB.C("test", "fs.files").FindOne(context.TODO(), bson.M{"filename": name}).Decode(&img)
	if err != nil {
		if name == " " {
			log.Printf("Failed to open %s: %v", name, err)
		} else {
			log.Printf("ProfileImage is emty")
		}
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	bucket, _ := gridfs.NewBucket(
		mongoDB.DB("test"),
	)
	var buf bytes.Buffer
	_, err = bucket.DownloadToStreamByName(
		img.Filename,
		&buf,
	)
	if err != nil {
		log.Println("GetProfileImg : Failed to open DownloadStream")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
	if err != nil {
		log.Printf("Failed to read downloadStream %s", err)
		http.Error(w, "Failed to read downloadStream", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, req, name, time.Now(), bytes.NewReader(buf.Bytes()))
}
