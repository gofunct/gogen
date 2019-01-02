package templates

import (
	"text/template"
)

func TemplateTemplate() *template.Template {
	return MustCreateTemplate("cmd", `package cmd

import (
	"github.com/izumin5210/clig/pkg/clib"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	{{- if .ViperEnabled}}
	"github.com/spf13/viper"
	{{- end}}
	"k8s.io/utils/exec"

	"{{.Package}}/pkg/{{.Name}}"
)
func NewDefault{{ToCamel .Name}}Command(wd clib.Path, build clib.Build) *cobra.Command {
	return New{{ToCamel .Name}}Command(&{{.Name}}.Ctx{
		WorkingDir: wd,
		IO:         clib.Stdio(),
		FS:         afero.NewOsFs(),
		{{- if .ViperEnabled}}
		Viper:      viper.New(),
		{{- end}}
		Exec:       exec.New(),
		Build:      build,
	})
}

func New{{ToCamel .Name}}Command(ctx *{{.Name}}.Ctx) *cobra.Command {
	cmd := &cobra.Command{
		Use: ctx.Build.AppName,
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			return errors.WithStack(ctx.Init())
		},
	}

	clib.AddLoggingFlags(cmd)

	cmd.AddCommand(
		clib.NewVersionCommand(ctx.IO, ctx.Build),
	)

	return cmd
}
`)
}
