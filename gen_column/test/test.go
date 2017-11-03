package test

type First struct {
	ID   int    `bson:"id,someothertags"`
	Name string `bson:"name"`
}
