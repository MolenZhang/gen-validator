package main

const (
	MgoTemplate = `
package {{.PackageName}}

import (
	"generator/gen_mgo/db"
	"sync"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	_CollectionName = "{{.CollectionName}}"
)

var (
	_C    *mgo.Collection
	_once sync.Once
)

type _Column struct {
	{{- range $i, $c := .Columns}}
	{{$c.Name}} string
	{{- end }}
}

var N _Column

func init() {
	{{- range $i, $c := .Columns}}
		N.{{$c.Name}} = "{{$c.Value}}"
	{{- end }}
}

{{- range $i, $c := .CustomStructs }}
type {{$c.Name}} {{$c.RawString}}
{{- end }}

type D {{ .StructString }}

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
`
)
