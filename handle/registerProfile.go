package handle

import (
	"context"
	"log"
	"mulbitchorong-backend/user"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RegisterProfile 함수는 유저의 프로필 정보를 받아 프로필 데이터를 등록하는 핸들러입니다.
func RegisterProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	_, fh, err := req.FormFile("profile")
	uid := req.FormValue("uid")
	uid = uid[1:(len(uid) - 1)]

	profileName := "profile_" + uid + "_mulbitchorong" + filepath.Ext(fh.Filename)

	bucket, err := gridfs.NewBucket(
		mongoDB.DB("test"),
	)
	registerRes := new(Res)
	if req.Method == "PUT" || req.Method == "POST" {
		var img *ImageFile
		err := mongoDB.C("test", "fs.files").FindOne(context.TODO(), bson.M{"filename": profileName}).Decode(&img)
		if err == nil {
			if err := bucket.Delete(img.ID); err != nil {
				registerRes.IsSuccess = false
				log.Println("file id : ", img.ID)
				log.Println("RegisterProfile : ", err)
				renderer.JSON(w, http.StatusOK, registerRes)
				return
			}
		}
	} else {
		log.Println("Request Method is ", req.Method)
	}

	u := user.New()
	err = mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": uid}).Decode(&u)
	if err != nil {
		registerRes.IsSuccess = false
		renderer.JSON(w, http.StatusOK, registerRes)
		log.Println("Not Found User : ", uid)
		log.Println("RegisterProfile : ", err)
		return
	}
	_, err = mongoDB.C("test", "users").UpdateOne(context.TODO(), bson.M{"_id": u.ID}, bson.M{"$set": bson.M{"profile_img": profileName}})
	if err != nil {
		registerRes.IsSuccess = false
		renderer.JSON(w, http.StatusOK, registerRes)
		log.Println("Failed to Update User Profile : ", profileName)
		return
	}
	file, _ := fh.Open()
	fe := filepath.Ext(fh.Filename)
	fileExt := fe[1:]
	opts := options.GridFSUpload().SetMetadata(bson.M{"uid": uid, "ext": fileExt})
	uploadStream, err := bucket.OpenUploadStream(
		profileName,
		opts,
	)
	if err != nil {
		registerRes.IsSuccess = false
		renderer.JSON(w, http.StatusOK, registerRes)
		log.Println("RegisterProfile : Failed to open UploadStream")
		return
	}
	if err := WriteToGridFile(file, uploadStream); err != nil {
		registerRes.IsSuccess = false
		renderer.JSON(w, http.StatusOK, registerRes)
		log.Println("Failed to Write File")
		return
	}
	registerRes.IsSuccess = true

	renderer.JSON(w, http.StatusCreated, registerRes)
}
