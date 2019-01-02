// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/gofunct/common/bingen/tool"
	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/excmd"
	"github.com/gofunct/gogen/pkg/gogencmd"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/usecase"
	"github.com/gofunct/gogen/pkg/protoc"
)

// Injectors from wire.go:

func NewUI(ctx *gogencmd.Ctx) cli.UI {
	io := gogencmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	return ui
}

func NewCommandExecutor(ctx *gogencmd.Ctx) excmd.Executor {
	io := gogencmd.ProvideIO(ctx)
	executor := excmd.NewExecutor(io)
	return executor
}

func NewGenerator(ctx *gogencmd.Ctx) module.Generator {
	io := gogencmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	generator := ProvideGenerator(ctx, ui)
	return generator
}

func NewScriptLoader(ctx *gogencmd.Ctx) module.ScriptLoader {
	io := gogencmd.ProvideIO(ctx)
	executor := excmd.NewExecutor(io)
	scriptLoader := ProvideScriptLoader(ctx, executor)
	return scriptLoader
}

func NewToolRepository(ctx *gogencmd.Ctx) (tool.Repository, error) {
	fs := gogencmd.ProvideFS(ctx)
	execInterface := gogencmd.ProvideExecer(ctx)
	io := gogencmd.ProvideIO(ctx)
	rootDir := gogencmd.ProvideRootDir(ctx)
	config := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	repository, err := protoc.ProvideToolRepository(config)
	if err != nil {
		return nil, err
	}
	return repository, nil
}

func NewProtocWrapper(ctx *gogencmd.Ctx) (protoc.Wrapper, error) {
	config := gogencmd.ProvideProtocConfig(ctx)
	fs := gogencmd.ProvideFS(ctx)
	execInterface := gogencmd.ProvideExecer(ctx)
	io := gogencmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	rootDir := gogencmd.ProvideRootDir(ctx)
	bingenConfig := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	repository, err := protoc.ProvideToolRepository(bingenConfig)
	if err != nil {
		return nil, err
	}
	wrapper := protoc.NewWrapper(config, fs, execInterface, ui, repository, rootDir)
	return wrapper, nil
}

func NewInitializeProjectUsecase(ctx *gogencmd.Ctx) usecase.InitializeProjectUsecase {
	fs := gogencmd.ProvideFS(ctx)
	execInterface := gogencmd.ProvideExecer(ctx)
	io := gogencmd.ProvideIO(ctx)
	rootDir := gogencmd.ProvideRootDir(ctx)
	config := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	ui := cli.UIInstance(io)
	generator := ProvideGenerator(ctx, ui)
	initializeProjectUsecase := ProvideInitializeProjectUsecase(ctx, config, ui, generator)
	return initializeProjectUsecase
}
