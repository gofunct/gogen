package context

import (
	"context"
	"github.com/gofunct/common/files"
	"github.com/gofunct/common/io"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"k8s.io/utils/exec"
	"log"
)

type Ctx struct {
	WorkingDir files.Path
	IO         io.IO
	FS         afero.Fs
	Exec       exec.Interface
	Build      Build
	Config     *Config
}

func (c *Ctx) GetRootDir() files.Path {
	return c.WorkingDir.Join(c.Build.AppName)
}

func (c *Ctx) GetIO() io.IO {
	return c.IO
}

func (c *Ctx) GetFS() afero.Fs {
	return c.FS
}

func (c *Ctx) GetExec() exec.Interface {
	return c.Exec
}

func (c *Ctx) GetBuild() Build {
	return c.Build
}

func (c *Ctx) GetImportPath() string {
	pkg, err := getImportPath(c.GetRootDir().String())
	if err != nil {
		log.Fatal("failed to get import path", err)
	}
	return pkg
}

func (c *Ctx) DefaultEntries(params *Params) []*Entry {
	//root := c.GetRootDir()
	return []*Entry{
		/*
		{Path: root.Join(".gitignore").String(), Template: templates.GitIgnoreTemplate()},
		{Path: root.Join(".reviewdog.yml").String(), Template: templates.ReviewDogTemplate()},
		{Path: root.Join(".travis.yml").String(), Template: templates.TravisTemplate()},
		{Path: root.Join("Makefile").String(), Template: templates.MakefileTemplate()},
		{Path: root.Join("cmd", params.Name, "main.go").String(), Template: templates.TemplateMain},
		{Path: root.Join("pkg", params.Name, "config.go").String(), Template: templates.ConfigTemplate()},
		{Path: root.Join("pkg", params.Name, "context.go").String(), Template: templates.ContextTemplate()},
		{Path: root.Join("pkg", params.Name, "cmd", "cmd.go").String(), Template: templates.TemplateTemplate()},
		*/
	}
}

func (c *Ctx) GetDefaultGexPkgs() []string {
	pkgs := []string{"github.com/mitchellh/gox"}
	pkgs = append(pkgs,
		"github.com/haya14busa/reviewdog/cmd/reviewdog",
		"github.com/kisielk/errcheck",
		"github.com/srvc/wraperr/cmd/wraperr",
		"golang.org/x/lint/golint",
		"honnef.co/go/tools/cmd/megacheck",
		"mvdan.cc/unparam",
	)
	return pkgs
}

func (c *Ctx) RunDepAndGex() error {
	ctx := context.TODO()
	run := c.GetRunFunc()

	err := run(ctx, "dep", "init")
	if err != nil {
		return err
	}

	if _, err := c.Exec.LookPath("gex"); err != nil {
		err = run(ctx, "go", "get", "github.com/izumin5210/gex/cmd/gex")
		if err != nil {
			return err
		}
	}

	pkgs := c.GetDefaultGexPkgs()

	gexArgs := make([]string, 2*len(pkgs))
	for i, pkg := range pkgs {
		gexArgs[2*i+0] = "--add"
		gexArgs[2*i+1] = pkg
	}

	err = run(ctx, "gex", gexArgs...)
	if err != nil {
		return err
	}

	return nil
}

func (c *Ctx) ParseEntries() error {
	params := c.NewParams()
	entries := c.DefaultEntries(params)

	for _, e := range entries {
		err := e.Create(c.FS, params)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Ctx) GetRunFunc() RunFunc {
	return func(ctx context.Context, program string, args ...string) error {
		cmd := c.Exec.CommandContext(ctx, program, args...)
		root := c.GetRootDir()

		cmd.SetStdin(c.IO.In())
		cmd.SetStdout(c.IO.Out())
		cmd.SetStderr(c.IO.Err())
		cmd.SetDir(root.String())
		zap.L().Debug("exec command", zap.String("cmd", c.Build.AppName), zap.Strings("args", args), zap.Stringer("dir", root))
		return cmd.Run()
	}
}
