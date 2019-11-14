package handle

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// GetImage 는 filename에 해당하는 이미지를 응답합니다.
func GetImage(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	name := ps.ByName("filename")

	gridfs := mongoDB.Session.DB("test").GridFS("fs")

	f, err := gridfs.Open(name)
	if err != nil {
		log.Printf("Failed to open %s: %v", name, err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	http.ServeContent(w, req, name, time.Now(), f) // Use proper last mod time
}
