package main

import (
	"bytes"
	"html/template"
)

func utilFuncMap() template.FuncMap {
	return template.FuncMap{
		"add":     add,
		"letters": letters,
	}
}

func add(a, b int) int {
	return a + b
}

func letters(count int, tag string) template.HTML {
	t, err := template.New("partail").Parse("<" + tag + ">{{.}}</" + tag + ">")

	if err != nil {
		logError(err.Error())
		return ""
	}

	buffer := bytes.Buffer{}

	for i := -1; i < count; i++ {
		if i < 0 {
			t.Execute(&buffer, nil)
		} else {
			t.Execute(&buffer, string(i+65))
		}
	}

	return template.HTML(buffer.String())
}
