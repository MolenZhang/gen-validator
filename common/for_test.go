package common

type StructOne struct {
	ID   int    `json:"id"`
	Name string `bson:"name"`
}

type StructTwo struct {
	Year     int
	Month    string
	Day      int
	One      StructOne
	OnePoint *StructOne
	OneArray []StructOne
}
