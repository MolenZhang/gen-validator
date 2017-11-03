package db

import (
	"fmt"
	"strings"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var (
	_Session *mgo.Session
	_DB      *mgo.Database
	_once    sync.Once
)

func InitMongoDB(addr, database, username, password string) {
	_once.Do(func() {
		info := &mgo.DialInfo{
			Addrs:    strings.Split(addr, ","),
			Timeout:  5 * time.Second,
			Database: database,
			Username: username,
			Password: password,
		}
		err := initMongoDBWith(info)
		if err != nil {
			panic(err)
		}
	})
}
func InitMongoDBWith(info *mgo.DialInfo) {
	_once.Do(func() {
		err := initMongoDBWith(info)
		if err != nil {
			panic(err)
		}
	})
}

func initMongoDBWith(info *mgo.DialInfo) error {
	if info == nil {
		return fmt.Errorf("DailInfo cannot be nil")
	}
	fmt.Printf("Start to connect db %v, db:%s, user:%s, pswd:%s\n", info.Addrs, info.Database, info.Username, info.Password)

	sess, err := mgo.DialWithInfo(info)
	if err != nil {
		return err
	}
	fmt.Printf("Success to connect db %v, db:%s, user:%s, pswd:%s\n", info.Addrs, info.Database, info.Username, info.Password)

	_Session = sess
	_DB = sess.DB(info.Database)

	return nil
}

func DB() *mgo.Database {
	if _DB == nil {
		panic("Forget to init mongodb ?")
	}
	return _DB
}

func Session() *mgo.Session {
	if _Session == nil {
		panic("Forget to init mongodb ?")
	}
	return _Session
}
