package templates

import "text/template"

func ConfigTemplate() *template.Template {
	return MustCreateTemplate("config", `package {{.Name}}

type Config struct {
}
`)
}
