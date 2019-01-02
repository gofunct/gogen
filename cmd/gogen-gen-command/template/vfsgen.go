// +build ignore

package template

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/izumin5210/grapi/cmd/grapi-gen-command/template"
)

func main() {
	err := vfsgen.Generate(template.FS, vfsgen.Options{
		PackageName:  "template",
		BuildTags:    "!vfsgen",
		VariableName: "FS",
	})
	if err != nil {
		log.Fatalln(err)
	}
}