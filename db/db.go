package db

import (
	"github.com/scriptonist/gowebapp/model"
	mgo "gopkg.in/mgo.v2"
)

// Config --
type Config struct {
	ConnectionString string
	DBName           string
}

// InitDB --
func InitDB(cfg Config) (*Mongo, error) {
	session, err := mgo.Dial(cfg.ConnectionString)
	if err != nil {
		return nil, err
	}
	return &Mongo{
		conn: session,
		DB:   session.DB(cfg.DBName),
	}, nil
}

// Mongo --
type Mongo struct {
	conn *mgo.Session
	DB   *mgo.Database
}

// SelectPeople --
func (m *Mongo) SelectPeople() ([]*model.Person, error) {
	return nil, nil
}
