package main

import (
	"bytes"
	"github.com/k0kubun/pp"
	"golang.org/x/tools/imports"
	"log"
	"os"
	"text/template"
)

// FormatSource is gofmt with addition of removing any unused imports.
// SEE: https://github.com/webrpc/webrpc/blob/v0.6.0/gen/golang/helpers.go
func FormatSource(source []byte) ([]byte, error) {
	return imports.Process("", source, &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 8,
	})
}

func main() {
	funcMap := template.FuncMap{
		"GetGreet": GetGreet,
	}

	tmpl, err := template.New("sample").
		Funcs(funcMap).
		ParseGlob("gen/golang/templates/*.go.tmpl")

	if err != nil {
		log.Fatal(err)
	}

	// generate the template
	genBuf := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(genBuf, "main.go.tmpl", struct{}{})
	if err != nil {
		pp.Println(err)
	}

	src, err := FormatSource(genBuf.Bytes())
	if err != nil {
		pp.Println(err)
	}

	err = os.WriteFile("out/main.go", src, 0644)
	if err != nil {
		pp.Println(err)
	}
}

func GetGreet() string {
	return "Hello world!"
}
