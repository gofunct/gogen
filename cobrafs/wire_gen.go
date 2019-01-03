// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cobrafs

import (
	"github.com/gofunct/common/ui"
	"github.com/gofunct/gogen/gogen"
)

// Injectors from injectors.go:

func newApp(command *Command) (*App, error) {
	ctx := ProvideCtx(command)
	gogenCtx := ProvideGogenCtx(ctx)
	fs := gogen.ProvideFS(gogenCtx)
	io := gogen.ProvideIO(gogenCtx)
	uiUI := ui.NewUI(io)
	rootDir := gogen.ProvideRootDir(gogenCtx)
	fileSystem := ProvideTemplateFS(command)
	shouldRunFunc := ProvideShouldRun(command)
	generator := NewGenerator(fs, uiUI, rootDir, fileSystem, shouldRunFunc)
	app := &App{
		Generator: generator,
		UI:        uiUI,
	}
	return app, nil
}
