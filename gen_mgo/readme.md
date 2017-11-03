# mgo的自动代码生成器
包装 mgo.v2版本，为每一个Collection都生成一套操作的代码.
减少业务代码的数量. 生成的api所有的错误（除了not found）都会panic,可以方便的统一处理数据库错误.

# 依赖
gen_mgo依赖 mgo.v2版本. 并且依赖全局的 gofmt 和 goreturns来格式化代码.
```
go get -u -v gopkg.in/mgo.v2
go get -u -v sourcegraph.com/sqs/goreturns

# mgo 也可以使用 godep 等来版本管理工具来放入 vendor.
```

# 使用
需要使用gen_mgo提供的初始化数据的方法
```
import (
	"generator/gen_mgo/db"
)

func main() {
	db.InitMongoDB("127.0.0.1", "test", "username", "password")
}
```
使用 go generate来生成代码
```
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
```
go generate 会自动在同级目录下创建一个名为 usercore 的目录，并且自动生成相关代码。
具体见 example
```
go generate ./...

# 安装依赖代码，例如 mgo
godep ensure -v

go build

```


# example

```
go get -u -v gopkg.in/mgo.v2
go get -u -v sourcegraph.com/sqs/goreturns

go run *.go -type UserCore -c user_core -path ./example

# start a mongodb
sudo docker run --rm -p 27017:27017 docker.ifeng.com/library/mongo:3.4.9

cd ./example
go run main.go
```