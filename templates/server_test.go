package templates

import "text/template"

func ServerTestTemplate() *template.Template {
	return MustCreateTemplate("server_test.go", serverTest)
}

var serverTest = `
package {{.Go.Package }}
{{if .Methods}}
import (
	"context"
	"testing"
{{range .Go.TestImports}}
	"{{.}}"
{{- end}}

	{{.PbGo.PackageName}} "{{ .PbGo.PackagePath }}"
)
{{$go := .Go -}}
{{$pbGo := .PbGo -}}
{{- range .Methods}}
func Test_{{$go.ServerName}}_{{.Method}}(t *testing.T) {
	svr := New{{$go.ServerName}}()

	ctx := context.Background()
	req := &{{.RequestGo $pbGo.PackageName}}{}

	resp, err := svr.{{.Method}}(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
{{end -}}
{{end -}}
`
