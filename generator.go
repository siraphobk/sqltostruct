package main

import (
	"bytes"
	"text/template"
)

type generator struct{}

func genStruct(t table) string {
	var result string

	structHeadTmplStr := "type {{.TableName}} struct {"

	tmpl, err := template.New("tablegen").Parse(structHeadTmplStr)
	if err != nil {
		panic(err.Error())
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, t); err != nil {
		panic(err.Error())
	}
	result += buf.String() + "\n"

	for _, column := range t.Columns {
		// TODO: Gen column
	}

	return result
}
