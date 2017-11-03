package test

type _FirstColumn struct {
	ID   string
	Name string
}

var FirstColumns _FirstColumn

func init() {
	FirstColumns.ID = "id"
	FirstColumns.Name = "name"
}
