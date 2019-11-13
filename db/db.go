package db

import "gopkg.in/mgo.v2"

// MongoDB 구조체는 몽고DB에 대한 정보를 담는 구조체입니다.
type MongoDB struct {
	mongoSession *mgo.Session
}

// NewMongoDB 메소드는 MongoDB 객체와 MongoDB Session을 생성합니다.
func NewMongoDB(addr string) (bool, *MongoDB) {
	s, err := mgo.Dial(addr)

	if err != nil {
		return false, nil
	}

	m := &MongoDB{mongoSession: s}
	return true, m
}
