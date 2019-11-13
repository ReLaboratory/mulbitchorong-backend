package handle

import (
	"mulbitchorong-backend/db"

	"github.com/unrolled/render"
)

var (
	renderer *render.Render
	mongoDB  *db.MongoDB
)

func init() {
	renderer = render.New()
}

// InitMongo 는 몽고DB의 초기 설정을 하는 함수입니다.
func InitMongo(addr string) error {
	m, err := db.NewMongoDB(addr)
	if err != nil {
		return err
	}
	mongoDB = m

	return nil
}
