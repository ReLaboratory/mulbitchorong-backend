package handle

import (
	"bufio"
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// Upload 함수는 유저 정보와 사용자 정보를 받아 이미지 업로드 기능을 수행하는 핸들러입니다.
func Upload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	for _, fileHeaders := range req.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			gridFile, err := mongoDB.DB.GridFS("fs").Create(fileHeader.Filename)
			if err != nil {
				//errorResponse(w, err, http.StatusInternalServerError)
				return
			}
			// gridFile.SetMeta()
			gridFile.SetName(fileHeader.Filename)
			if err := writeToGridFile(file, gridFile); err != nil {
				//errorResponse(w, err, http.StatusInternalServerError)
				return
			}
		}
	}
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
