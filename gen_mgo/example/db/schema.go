package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SubSub struct {
	Bool bool
	Int  int
}

type SubDoc struct {
	Age    int
	Name   string `bson:""`
	SubSub SubSub
}

//go:generate gen_mgo -type UserCore -c user_core
type UserCore struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	GUID       uint64        `bson:"guid"`
	Username   string        `bson:"username"`
	Mail       string        `bson:"mail"`
	Mobile     string        `bson:"mobile"`
	Password   string        `bson:"password"`
	Nickname   string        `bson:"nickname"`
	Image      string        `bson:"image"`
	UserStatus int8          `bson:"user_status"`
	CreateTime time.Time     `bson:"create_time"`
	UpdateTime time.Time     `bson:"update_time"`
	Sub        SubDoc
}
