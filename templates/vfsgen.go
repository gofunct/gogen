// +build ignore

package main

import (
	"github.com/gofunct/gogen/templates"
	"log"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(templates.VirtualFS, vfsgen.Options{
		PackageName:  "templates",
		BuildTags:    "!vfsgen",
		VariableName: "VirtualFS",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
