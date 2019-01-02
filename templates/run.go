package templates

import "text/template"

func RunTemplate() *template.Template {
	return MustCreateTemplate("run.go", run)
}

var run = `package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	fmt.Println("It works!")
	return 0
}
`
