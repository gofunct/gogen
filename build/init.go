package build

import (
	"bytes"
	"context"
	"errors"
	"github.com/gofunct/common/templates"
	"go/build"
	"go/format"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	)

var (
	BuildContext = build.Default
)

func newInitCommand(c *Ctx) *cobra.Command {
	var (
		skipViper bool
	)

	cmd := &cobra.Command{
		Use:  "init",
		Args: cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			name := args[0]
			root := c.WorkingDir.Join(name)
			pkg, err := getImportPath(root.String())
			if err != nil {
				return err
			}

			params := struct {
				Name         string
				Package      string
				ViperEnabled bool
			}{Name: name, Package: pkg, ViperEnabled: !skipViper}

			entries := []*entry{
				{Path: root.Join(".gitignore").String(), Template: templates.GitIgnoreTemplate()},
				{Path: root.Join(".reviewdog.yml").String(), Template: templates.ReviewDogTemplate()},
				{Path: root.Join(".travis.yml").String(), Template: templates.TravisTemplate()},
				{Path: root.Join("Makefile").String(), Template: templates.MakefileTemplate()},
				{Path: root.Join("cmd", params.Name, "main.go").String(), Template: templates.TemplateMain},
				{Path: root.Join("pkg", params.Name, "config.go").String(), Template: templates.ConfigTemplate(), Skipped: skipViper},
				{Path: root.Join("pkg", params.Name, "context.go").String(), Template: templates.ContextTemplate()},
				{Path: root.Join("pkg", params.Name, "cmd", "cmd.go").String(), Template: templates.TemplateTemplate()},
			}

			for _, e := range entries {
				if e.Skipped {
					continue
				}
				err = e.Create(c.FS, params)
				if err != nil {
					return err
				}
			}

			ctx := context.TODO()

			run := func(ctx context.Context, name string, args ...string) error {
				cmd := c.Exec.CommandContext(ctx, name, args...)
				cmd.SetStdin(c.IO.In())
				cmd.SetStdout(c.IO.Out())
				cmd.SetStderr(c.IO.Err())
				cmd.SetDir(root.String())
				zap.L().Debug("exec command", zap.String("cmd", name), zap.Strings("args", args), zap.Stringer("dir", root))
				return cmd.Run()
			}

			err = run(ctx, "dep", "init")
			if err != nil {
				return err
			}

			if _, err := c.Exec.LookPath("gex"); err != nil {
				err = run(ctx, "go", "get", "github.com/izumin5210/gex/cmd/gex")
				if err != nil {
					return err
				}
			}

			pkgs := []string{"github.com/mitchellh/gox"}
			pkgs = append(pkgs,
				"github.com/haya14busa/reviewdog/cmd/reviewdog",
				"github.com/kisielk/errcheck",
				"github.com/srvc/wraperr/cmd/wraperr",
				"golang.org/x/lint/golint",
				"honnef.co/go/tools/cmd/megacheck",
				"mvdan.cc/unparam",
			)
			bingenArgs := make([]string, 2*len(pkgs))
			for i, pkg := range pkgs {
				bingenArgs[2*i+0] = "--add"
				bingenArgs[2*i+1] = pkg
			}

			err = run(ctx, "bingen", bingenArgs...)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&skipViper, "skip-viper", false, "Do not use viper")

	return cmd
}

func getImportPath(rootPath string) (importPath string, err error) {
	for _, gopath := range filepath.SplitList(BuildContext.GOPATH) {
		prefix := filepath.Join(gopath, "src") + string(filepath.Separator)
		// FIXME: should not use strings.HasPrefix
		if strings.HasPrefix(rootPath, prefix) {
			importPath = filepath.ToSlash(strings.Replace(rootPath, prefix, "", 1))
			break
		}
	}
	if importPath == "" {
		err = errors.New("failed to get the import path")
	}
	return
}

type entry struct {
	Template *template.Template
	Path     string
	Skipped  bool
}

func (e *entry) Create(fs afero.Fs, params interface{}) error {
	dir := filepath.Dir(e.Path)
	if ok, err := afero.DirExists(fs, dir); err != nil {
		return err
	} else if !ok {
		zap.L().Debug("create a directory", zap.String("dir", dir))
		err = fs.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	buf := new(bytes.Buffer)
	err := e.Template.Execute(buf, params)
	if err != nil {
		return err
	}

	data := buf.Bytes()

	if filepath.Ext(e.Path) == ".go" {
		data, err = format.Source(data)
		if err != nil {
			return err
		}
	}

	zap.L().Debug("create a new file", zap.String("path", e.Path))
	err = afero.WriteFile(fs, e.Path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
