package db

import "gopkg.in/mgo.v2"

// MongoDB 구조체는 몽고DB에 대한 정보를 담는 구조체입니다.
type MongoDB struct {
	Session *mgo.Session
}

// NewMongoDB 메서드는 MongoDB 객체와 MongoDB Session을 생성합니다.
func NewMongoDB(addr string) (*MongoDB, error) {
	s, err := mgo.Dial(addr)

	if err != nil {
		return nil, err
	}

	m := &MongoDB{}
	m.Session = s
	return m, nil
}
