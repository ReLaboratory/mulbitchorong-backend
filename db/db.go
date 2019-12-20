package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB 구조체는 몽고DB에 대한 정보를 담는 구조체입니다.
type MongoDB struct {
	Client      *mongo.Client
	Ctx         context.Context
	Cancel      context.CancelFunc
	databases   map[string]*mongo.Database
	collections map[string]map[string]*mongo.Collection
}

// C 메서드는 해당하는 DB의 컬렉션을 반환합니다.
func (m *MongoDB) C(db, c string) *mongo.Collection {
	if m.collections == nil {
		m.initCollectionMap()
	}
	if m.collections[db] == nil {
		m.collections[db] = make(map[string]*mongo.Collection)
	}
	_, exists := m.collections[db][c]
	if exists {
		m.AddCollection(db, c)
	}
	return m.collections[db][c]
}

func (m *MongoDB) initCollectionMap() {
	m.collections = make(map[string]map[string]*mongo.Collection)
}

// DB 메서드는 해당하는 DB를 반환합니다.
func (m *MongoDB) DB(db string) *mongo.Database {
	if m.databases == nil {
		m.AddDatabase(db)
	}
	return m.databases[db]
}

// AddCollection 메서드는 Collection Map을 통해 사용할 Collection을 추가합니다.
func (m *MongoDB) AddCollection(db, c string) {
	if m.collections == nil {
		m.initCollectionMap()
	}
	if m.collections[db] == nil {
		m.collections[db] = make(map[string]*mongo.Collection)
	}
	m.collections[db][c] = m.databases[db].Collection(c)
}

// SetCollections 메서드는 특정 DB에서 사용할 Collection을 설정합니다.
func (m *MongoDB) SetCollections(db string, cs []string) {
	for i := 0; i < len(cs); i++ {
		m.AddCollection(db, cs[i])
	}
}

// AddDatabase 메서드는 Database Map에 사용할 DB를 추가합니다
func (m *MongoDB) AddDatabase(db string) {
	if m.databases == nil {
		m.databases = make(map[string]*mongo.Database)
	}
	if m.Client == nil {
		log.Println("Client is Nil")
	}
	m.databases[db] = m.Client.Database(db)
}

// SetDatabases 메서드는 특정 DB에서 사용할 Collection을 설정합니다.
func (m *MongoDB) SetDatabases(dbs []string) {
	for i := 0; i < len(dbs); i++ {
		m.AddDatabase(dbs[i])
	}
}

// Connect 메서드는 몽고DB에 연결합니다.
func (m *MongoDB) Connect() error {
	if err := m.Client.Connect(m.Ctx); err != nil {
		return err
	}
	return nil
}

// Disconnect 메서드는 몽고DB 연결을 끊습니다.
func (m *MongoDB) Disconnect() {
	m.Client.Disconnect(m.Ctx)
}

// Ping 메서드는 몽고DB 연결을 확인합니다.
func (m *MongoDB) Ping() error {
	if err := m.Client.Ping(m.Ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// NewClient 메서드는 몽고DB Client를 생성합니다.
func (m *MongoDB) NewClient(addr string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		return err
	}
	m.Client = client
	return nil
}

// NewMongoDB 메서드는 MongoDB 객체와 MongoDB Session을 생성합니다.
func NewMongoDB(addr string, dbs []string) (*MongoDB, error) {
	m := &MongoDB{}
	if err := m.NewClient(addr); err != nil {
		log.Println("MongoDB-NewClient 실패 ", err)
		return nil, err
	}

	for i := 0; i < len(dbs); i++ {
		m.AddDatabase(dbs[i])
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m.Ctx = ctx
	m.Cancel = cancel

	if err := m.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := m.Ping(); err != nil {
		log.Println("DB : Failed to Connect DB")
		log.Fatal(err)
		return nil, err
	}
	return m, nil
}
