//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/izumin5210/gex/pkg/tool"

	"github.com/gofunct/common/excmd"
	"github.com/gofunct/gogen/gogen"
	"github.com/gofunct/gogen/module"
	"github.com/gofunct/gogen/protoc"
	"github.com/gofunct/gogen/usecase"
)

func NewUI(*gogen.Ctx) cli.UI {
	wire.Build(Set)
	return nil
}

func NewCommandExecutor(*gogen.Ctx) excmd.Executor {
	wire.Build(Set)
	return nil
}

func NewGenerator(*gogen.Ctx) module.Generator {
	wire.Build(Set)
	return nil
}

func NewScriptLoader(*gogen.Ctx) module.ScriptLoader {
	wire.Build(Set)
	return nil
}

func NewToolRepository(*gogen.Ctx) (tool.Repository, error) {
	wire.Build(Set)
	return nil, nil
}

func NewProtocWrapper(*gogen.Ctx) (protoc.Wrapper, error) {
	wire.Build(Set)
	return nil, nil
}

func NewInitializeProjectUsecase(*gogen.Ctx) usecase.InitializeProjectUsecase {
	wire.Build(Set)
	return nil
}
