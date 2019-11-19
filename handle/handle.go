package handle

import (
	"fmt"
	"mulbitchorong-backend/db"
	"os"

	"github.com/unrolled/render"
)

// ImageFile 은 이미지 파일에 대한 정보를 담고 있는 구조체입니다.
type ImageFile struct {
	Filename string   `bson:"filename"`
	MetaData FileMeta `bson:"metadata"`
}

// ImageName 은 이미지 파일의 이름을 담고 있는 구조체입니다.
type ImageName struct {
	Name string `json:"imgName"`
}

// ImageUploader 는 이미지 파일을 업로드한 유저의 ID를 담고 있는 구조체입니다.
type ImageUploader struct {
	UID string `json:"uid"`
}

// FileMeta 는 Upload할 파일의 메타정보를 담는 구조체입니다.
type FileMeta struct {
	Inode int
	UID   string `bson:"uid" json:"uid"`
	Ext   string `bson:"ext" json:"ext"`
}

// Res 는 Default 응답값인 성공여부를 담는 구조체입니다.
type Res struct {
	IsSuccess bool `json:"isSuccess"`
}

var (
	renderer *render.Render
	mongoDB  *db.MongoDB
)

func init() {
	renderer = render.New()
}

// InitMongo 는 몽고DB의 초기 설정을 하는 함수입니다.
func InitMongo(addr string) error {
	var dbID, dbPw string
	fi, err := os.Open("db_account.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fmt.Fscan(fi, &dbID, &dbPw)

	m, err := db.NewMongoDB(addr)
	if err != nil {
		return err
	}
	mongoDB = m
	if err := mongoDB.Session.DB("admin").Login(dbID, dbPw); err != nil {
		panic(err)
	}
	return nil
}
