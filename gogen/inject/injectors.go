//+build wireinject

package inject

import (
	"github.com/gofunct/common/ui"
	"github.com/google/wire"
	"github.com/izumin5210/gex/pkg/tool"

	"github.com/gofunct/common/executor"
	"github.com/gofunct/gogen/gogen"
	"github.com/gofunct/gogen/module"
	"github.com/gofunct/gogen/protoc"
	"github.com/gofunct/gogen/usecase"
)

func NewUI(*gogen.Ctx) ui.UI {
	wire.Build(InjectorSet)
	return nil
}

func NewCommandExecutor(*gogen.Ctx) executor.Executor {
	wire.Build(InjectorSet)
	return nil
}

func NewGenerator(*gogen.Ctx) module.Generator {
	wire.Build(InjectorSet)
	return nil
}

func NewScriptLoader(*gogen.Ctx) module.ScriptLoader {
	wire.Build(InjectorSet)
	return nil
}

func NewToolRepository(*gogen.Ctx) (tool.Repository, error) {
	wire.Build(InjectorSet)
	return nil, nil
}

func NewProtocWrapper(*gogen.Ctx) (protoc.Wrapper, error) {
	wire.Build(InjectorSet)
	return nil, nil
}

func NewInitializeProjectUsecase(*gogen.Ctx) usecase.InitializeProjectUsecase {
	wire.Build(InjectorSet)
	return nil
}
