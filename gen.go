package abi

import (
	"bytes"
	"strings"
	"fmt"
	"go/format"
	"text/template"
)

const tmplSourceGo = `
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

// Credit to the go-ethereum authors for pioneering work on the Solidity abigen, on which this work is based.

package {{.Package}}

{{$structs := .Structs}}
{{range $structs}}
	// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
	type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
	}
{{end}}`

type tmplData struct {
	Package string
	Structs map[string]*tmplStruct
}

type tmplStruct struct {
	Name string
	Fields []*tmplField
}

type tmplField struct {
	Type    string   // Field type representation depends on target binding language
	Name    string   // Field name converted from the raw user-defined field name
	// SolKind abi.Type // Raw abi type information // TODO
}

func takeLast(s string) string {
	str := strings.Split(s, "::")
	var a string
	for _, x := range str {
		a = a + strings.ToUpper(string(x[0])) + x[1:]
	}
	return a
}

func Generate(pkg string, abi *ABI) (string, error) {
	buffer := new(bytes.Buffer)
	data := &tmplData{
		Package: pkg,
		Structs: make(map[string]*tmplStruct, 0),
	}
	for _, aStruct := range abi.Structs {
		resultStruct := &tmplStruct{}
		resultStruct.Name = "Struct"+takeLast(aStruct.Name)
		resultStruct.Fields = make([]*tmplField, 0)
		for _, field := range aStruct.Members {
			var _type string
			// TODO more types
			if field.Type == "core::integer::u128" {
				_type = "[2]uint64"	
			} else {
				_type = "Type"+takeLast(field.Type) 
			}
			resultStruct.Fields = append(resultStruct.Fields, &tmplField{
				Type: _type,
				Name: strings.ToUpper(string(field.Name[0]))+field.Name[1:],
			})
		}
		data.Structs[aStruct.Name] = resultStruct
	}
	tmpl := template.Must(template.New("").Parse(tmplSourceGo))
	if err := tmpl.Execute(buffer, data); err != nil {
		return "", err
	}
	code, err := format.Source(buffer.Bytes())
	if err != nil {
		return "", fmt.Errorf("%v\n%s", err, buffer)
	}
	return string(code), nil
}
