package common

import "testing"

func AssertStringEqual(t *testing.T, s1, s2 string) {
	if s1 != s2 {
		t.Errorf("%s should be equal to %s", s1, s2)
	}
}

func AssertIntEqual(t *testing.T, s1, s2 int) {
	if s1 != s2 {
		t.Errorf("%s should be equal to %s", s1, s2)
	}
}

func TestGetPackageFromFile(t *testing.T) {
	pkg, _ := GetPackageFromFile("./for_test.go")
	AssertStringEqual(t, pkg.Name, "common")
	AssertIntEqual(t, len(pkg.Structs), 2)
	AssertStringEqual(t, pkg.Structs[0].Name, "StructOne")
	AssertIntEqual(t, len(pkg.Structs[0].Columns), 2)
	AssertStringEqual(t, pkg.Structs[0].Columns[0].Name, "ID")
	AssertStringEqual(t, pkg.Structs[0].Columns[0].Type, "int")
	AssertStringEqual(t, pkg.Structs[0].Columns[0].Tag, "`json:\"id\"`")
	AssertStringEqual(t, pkg.Structs[0].Columns[1].Name, "Name")
	AssertStringEqual(t, pkg.Structs[0].Columns[1].Type, "string")
	AssertStringEqual(t, pkg.Structs[0].Columns[1].Tag, "`bson:\"name\"`")
	AssertStringEqual(t, pkg.Structs[1].Name, "StructTwo")
	AssertIntEqual(t, len(pkg.Structs[1].Columns), 6)
	AssertStringEqual(t, pkg.Structs[1].Columns[0].Name, "Year")
	AssertStringEqual(t, pkg.Structs[1].Columns[0].Tag, "")
	AssertStringEqual(t, pkg.Structs[1].Columns[0].Type, "int")
	AssertStringEqual(t, pkg.Structs[1].Columns[1].Name, "Month")
	AssertStringEqual(t, pkg.Structs[1].Columns[1].Tag, "")
	AssertStringEqual(t, pkg.Structs[1].Columns[1].Type, "string")
	AssertStringEqual(t, pkg.Structs[1].Columns[2].Name, "Day")
	AssertStringEqual(t, pkg.Structs[1].Columns[2].Tag, "")
	AssertStringEqual(t, pkg.Structs[1].Columns[2].Type, "int")

	AssertStringEqual(t, pkg.Structs[1].Columns[3].Name, "One")
	AssertStringEqual(t, pkg.Structs[1].Columns[3].Tag, "")
	AssertStringEqual(t, pkg.Structs[1].Columns[3].Type, "StructOne")

	AssertStringEqual(t, pkg.Structs[1].Columns[4].Name, "OnePoint")
	AssertStringEqual(t, pkg.Structs[1].Columns[4].Tag, "")
	AssertStringEqual(t, pkg.Structs[1].Columns[4].Type, "") // TODO 不识别高级的类似指针，数组以及他们的组合
	AssertStringEqual(t, pkg.Structs[1].Columns[5].Name, "OneArray")
	AssertStringEqual(t, pkg.Structs[1].Columns[5].Tag, "")
	AssertStringEqual(t, pkg.Structs[1].Columns[5].Type, "array") // TODO 不识别高级的类似指针，数组以及他们的组合
}
