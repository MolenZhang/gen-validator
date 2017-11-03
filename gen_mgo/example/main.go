package main

import (
	"fmt"
	"generator/gen_mgo/db"
	"generator/gen_mgo/example/db/usercore"
)

func main() {
	db.InitMongoDB("127.0.0.1", "test", "", "")

	usercore.Insert(usercore.D{
		GUID:     12345678,
		Username: "Bob",
		Mail:     "ifeng@ifeng.com",
	})
	d := usercore.Find(nil).One()
	fmt.Println(d)
}
