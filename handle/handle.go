package handle

import (
	"context"
	"fmt"
	"log"
	"mulbitchorong-backend/db"
	"mulbitchorong-backend/user"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
)

// ImageFile 은 이미지 파일에 대한 정보를 담고 있는 구조체입니다.
type ImageFile struct {
	ID       primitive.ObjectID `bson:"_id"`
	Filename string             `bson:"filename"`
	MetaData FileMeta           `bson:"metadata"`
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
NEW_MONGO:
	db, err := db.NewMongoDB(
		"mongodb://mul2019re:bit2019re@cluster0-shard-00-00-zdtbn.mongodb.net:27017,cluster0-shard-00-01-zdtbn.mongodb.net:27017,cluster0-shard-00-02-zdtbn.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority",
		[]string{"test"},
	)
	if err != nil {
		log.Println("db.NewMongoDB : ", err)
		goto NEW_MONGO
	}
	InitMongo(db)
}

// InitMongo 는 몽고DB의 초기 설정을 하는 함수입니다.
func InitMongo(db *db.MongoDB) error {
	dbs := []string{
		"test",
		"uploadfile",
	}
	db.SetDatabases(dbs)
	testCs := []string{
		"users",
		"fs.files",
		"fs.chunks",
	}
	uploadCs := []string{
		"fs.files",
		"fs.chunks",
	}
	mongoDB = db
	mongoDB.SetCollections("test", testCs)
	mongoDB.SetCollections("uploadfile", uploadCs)
	/* TEST */

	IDCheck := user.User{}
	mongoDB.DB("test").Collection("users").FindOne(context.TODO(), bson.M{"uid": "yunseoh68"}).Decode(&IDCheck)
	mongoDB.C("test", "users").FindOne(context.TODO(), bson.M{"uid": "yunseoh68"}).Decode(&IDCheck)
	fmt.Println(IDCheck.Name)
	return nil
}
