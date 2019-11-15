package handle

import (
	"bufio"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Upload 함수는 유저 정보와 사용자 정보를 받아 이미지 업로드 기능을 수행하는 핸들러입니다.
func Upload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uploadRes := new(Res)

	req.ParseForm()
	_, fh, err := req.FormFile("files")
	iname := req.FormValue("imgName")
	uid := req.FormValue("uid")

	iname = iname[1:(len(iname) - 1)]
	uid = uid[1:(len(uid) - 1)]

	file, _ := fh.Open()
	timeNow := time.Now().Format("2006-01-02-15:04:05")
	uFileName := string(iname + "_" + timeNow + filepath.Ext(fh.Filename))
	gridFile, err := mongoDB.Session.DB("test").GridFS("fs").Create(uFileName)
	if err != nil {
		uploadRes.IsSuccess = false
	}
	fe := filepath.Ext(fh.Filename)
	fileExt := fe[1:]
	gridFile.SetMeta(bson.M{"uid": uid, "ext": fileExt})
	gridFile.SetName(uFileName)

	if err := WriteToGridFile(file, gridFile); err != nil {
		uploadRes.IsSuccess = false
	} else {
		uploadRes.IsSuccess = true
	}
	renderer.JSON(w, http.StatusCreated, uploadRes)
}

// WriteToGridFile 은 GridFile로 파일 쓰기를 수행합니다.
func WriteToGridFile(file multipart.File, gridFile *mgo.GridFile) error {
	reader := bufio.NewReader(file)
	defer func() { file.Close() }()
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return errors.New("Could not read the input file")
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := gridFile.Write(buf[:n]); err != nil {
			return errors.New("Could not write to GridFs for " + gridFile.Name())
		}
	}
	gridFile.Close()
	return nil
}
