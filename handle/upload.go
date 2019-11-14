package handle

import (
	"bufio"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FileMeta 는 Upload할 파일의 메타정보를 담는 구조체입니다.
type FileMeta struct {
	Inode int
	UID   string
}

type res struct {
	IsSuccess bool `json:"isSuccess"`
}

// Upload 함수는 유저 정보와 사용자 정보를 받아 이미지 업로드 기능을 수행하는 핸들러입니다.
func Upload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uploadRes := new(res)

	req.ParseForm()
	_, fh, err := req.FormFile("files")
	iname := req.FormValue("imgName")
	uid := req.FormValue("uid")

	file, _ := fh.Open()
	uFileName := string(iname + filepath.Ext(fh.Filename))
	gridFile, err := mongoDB.Session.DB("test").GridFS("fs").Create(uFileName)
	if err != nil {
		uploadRes.IsSuccess = false
	}
	gridFile.SetMeta(bson.M{"user": uid})
	gridFile.SetName(uFileName)

	if err := writeToGridFile(file, gridFile); err != nil {
		uploadRes.IsSuccess = false
	} else {
		uploadRes.IsSuccess = true
	}
	renderer.JSON(w, http.StatusCreated, uploadRes)
}

func writeToGridFile(file multipart.File, gridFile *mgo.GridFile) error {
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
