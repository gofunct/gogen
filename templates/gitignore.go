package templates

import "text/template"

func GitIgnoreTemplate() *template.Template {

	return MustCreateTemplate("gitignore", `/bin
/dist
/vendor
/.idea
`)
}
