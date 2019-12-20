package handle

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"mulbitchorong-backend/user"
	"net/http"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Upload 함수는 유저 정보와 사용자 정보를 받아 이미지 업로드 기능을 수행하는 핸들러입니다.
func Upload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	_, fh, err := req.FormFile("files")
	iname := req.FormValue("imgName")
	uid := req.FormValue("uid")
	iname = iname[1:(len(iname) - 1)]
	uid = uid[1:(len(uid) - 1)]

	result := user.New()
	uploadRes := new(Res)
	err = mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": uid}).Decode(&result)
	if err != nil {
		uploadRes.IsSuccess = false
		renderer.JSON(w, http.StatusOK, uploadRes)
		log.Println("Upload : Non-existent user ", uid)
		return
	}

	bucket, err := gridfs.NewBucket(
		mongoDB.DB("uploadfile"),
	)
	if err != nil {
		log.Println(err)
		uploadRes.IsSuccess = false
		renderer.JSON(w, http.StatusOK, uploadRes)
		log.Println("Upload : Failed to create Bucket")
		return
	}

	file, _ := fh.Open()
	timeNow := time.Now().Format("2006-01-02-15:04:05")
	uFileName := string(iname + "_" + timeNow + filepath.Ext(fh.Filename))
	fe := filepath.Ext(fh.Filename)
	fileExt := fe[1:]
	opts := options.GridFSUpload().SetMetadata(bson.M{"uid": uid, "ext": fileExt})
	uploadStream, err := bucket.OpenUploadStream(
		uFileName,
		opts,
	)
	if err != nil {
		log.Println(err)
	}

	if err := WriteToGridFile(file, uploadStream); err != nil {
		uploadRes.IsSuccess = false
		log.Println("Upload : Upload failed")
	} else {
		uploadRes.IsSuccess = true
		log.Println("Upload : Upload successful")
	}
	renderer.JSON(w, http.StatusCreated, uploadRes)
}

// WriteToGridFile 은 GridFile로 파일 쓰기를 수행합니다.
func WriteToGridFile(file multipart.File, uploadStream *gridfs.UploadStream) error {
	reader := bufio.NewReader(file)
	defer func() { file.Close() }()
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return errors.New("Could not read the input file")
		}
		if n == 0 {
			break
		}
		if _, err := uploadStream.Write(buf[:n]); err != nil {
			return errors.New("Could not write to GridFs")
		}
	}
	uploadStream.Close()
	return nil
}
