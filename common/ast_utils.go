package common

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func MustParseFile(filename string) *ast.File {
	fset := token.NewFileSet()
	var src interface{}
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		panic(err)
	}
	return f
}

func GetAllTypeMap(file *ast.File) map[string]*ast.TypeSpec {
	if file == nil {
		return nil
	}
	m := make(map[string]*ast.TypeSpec)
	for _, f := range file.Decls {
		gen, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		if gen.Tok != token.TYPE {
			continue
		}

		for _, v := range gen.Specs {
			spec := v.(*ast.TypeSpec)
			name := spec.Name.Name
			m[name] = spec
		}
	}
	return m
}

func GetAllTypeMapFromFile(filename string) map[string]*ast.TypeSpec {
	return GetAllTypeMap(MustParseFile(filename))
}

func GetAllStructMapFromFile(filename string) map[string]*ast.StructType {
	return GetAllStructMap(MustParseFile(filename))
}

func GetAllStructMap(file *ast.File) map[string]*ast.StructType {
	if file == nil {
		return nil
	}
	m := make(map[string]*ast.StructType)
	for _, f := range file.Decls {
		gen, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		if gen.Tok != token.TYPE {
			continue
		}

		for _, v := range gen.Specs {
			spec := v.(*ast.TypeSpec)
			name := spec.Name.Name
			if s, ok := spec.Type.(*ast.StructType); ok {
				m[name] = s
			}
		}
	}
	return m
}

func GetStructFromFileByName(filename, name string) *ast.StructType {
	return GetStructByName(MustParseFile(filename), name)
}

func GetStructByName(file *ast.File, name string) *ast.StructType {
	all := GetAllStructMap(file)
	if all == nil {
		return nil
	}
	return all[name]
}

func GetPackageName(file *ast.File) string {
	if file == nil {
		return ""
	}
	return file.Name.Name
}

func GetRawBytesWithPosEnd(filename string, pos, end token.Pos) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Seek(int64(pos)-1, 0)
	size := int(end) - int(pos)
	buf := make([]byte, size)
	_, err = f.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf
}
