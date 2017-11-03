package usercore

import (
	"generator/gen_mgo/db"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	_CollectionName = "user_core"
)

var (
	_C    *mgo.Collection
	_once sync.Once
)

type _Column struct {
	ID              string
	GUID            string
	Username        string
	Mail            string
	Mobile          string
	Password        string
	Nickname        string
	Image           string
	UserStatus      string
	CreateTime      string
	UpdateTime      string
	Sub             string
	Sub_Age         string
	Sub_Name        string
	Sub_SubSub      string
	Sub_SubSub_Bool string
	Sub_SubSub_Int  string
}

var N _Column

func init() {
	N.ID = "_id"
	N.GUID = "guid"
	N.Username = "username"
	N.Mail = "mail"
	N.Mobile = "mobile"
	N.Password = "password"
	N.Nickname = "nickname"
	N.Image = "image"
	N.UserStatus = "user_status"
	N.CreateTime = "create_time"
	N.UpdateTime = "update_time"
	N.Sub = "sub"
	N.Sub_Age = "sub.age"
	N.Sub_Name = "sub.name"
	N.Sub_SubSub = "sub.subsub"
	N.Sub_SubSub_Bool = "sub.subsub.bool"
	N.Sub_SubSub_Int = "sub.subsub.int"
}

type SubDoc struct {
	Age    int
	Name   string `bson:""`
	SubSub SubSub
}
type SubSub struct {
	Bool bool
	Int  int
}

type D struct {
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

type Query struct {
	Query *mgo.Query
}

func (q *Query) Select(selector interface{}) *Query {
	q.Query = q.Query.Select(selector)
	return q
}

func (q *Query) Skip(n int) *Query {
	q.Query = q.Query.Skip(n)
	return q
}

func (q *Query) Count(n int) int {
	n, err := q.Query.Count()
	if err == mgo.ErrNotFound {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return n
}

func (q *Query) All(n int) []D {
	var results []D
	err := q.Query.All(&results)
	if err == mgo.ErrNotFound {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return results
}

func (q *Query) One() *D {
	var result D
	err := q.Query.One(&result)
	if err == mgo.ErrNotFound {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

func C() *mgo.Collection {
	if _C == nil {
		_once.Do(func() {
			_C = db.DB().C(_CollectionName)
		})
	}
	return _C
}

func Find(query interface{}) *Query {
	return &Query{Query: C().Find(query)}
}

func FindId(id interface{}) *Query {
	return &Query{Query: C().FindId(id)}
}

func Insert(docs ...interface{}) {
	err := C().Insert(docs...)
	if err != nil {
		panic(err)
	}
}

// 1 if remove ok, 0 if not found. panic if any error
func Remove(selector interface{}) int {
	err := C().Remove(selector)
	if err == mgo.ErrNotFound {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return 1
}

// return removed number
func RemoveAll(selector interface{}) int {
	info, err := C().RemoveAll(selector)
	if err == mgo.ErrNotFound {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return info.Removed
}

func RemoveId(id interface{}) int {
	err := C().RemoveId(id)
	if err == mgo.ErrNotFound {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return 1
}

// 1 if update ok, 0 if not found. panic if any error
func Update(selector interface{}, update interface{}) int {
	err := C().Update(selector, update)
	if err == mgo.ErrNotFound {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return 1
}

// return updated number
func UpdateAll(selector interface{}, update interface{}) int {
	info, err := C().UpdateAll(selector, update)
	if err == mgo.ErrNotFound {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return info.Updated
}

// 1 if update ok, 0 if not found. panic if any error
func UpdateId(id interface{}, update interface{}) int {
	err := C().UpdateId(id, update)
	if err == mgo.ErrNotFound {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return 1
}

// return upserted _id
func Upsert(selector interface{}, update interface{}) interface{} {
	info, err := C().Upsert(selector, update)
	if err == mgo.ErrNotFound {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return info.UpsertedId
}
