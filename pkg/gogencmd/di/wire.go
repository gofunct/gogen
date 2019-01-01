//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/gofunct/common/bingen/tool"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/excmd"
	"github.com/gofunct/gogen/grapicmd"
	"github.com/gofunct/gogen/pkg/gogencmdinternal/module"
	"github.com/gofunct/gogen/pkg/gogencmdinternal/usecase"
	"github.com/gofunct/gogen/protoc"
)

func NewUI(*grapicmd.Ctx) cli.UI {
	wire.Build(Set)
	return nil
}

func NewCommandExecutor(*grapicmd.Ctx) excmd.Executor {
	wire.Build(Set)
	return nil
}

func NewGenerator(*grapicmd.Ctx) module.Generator {
	wire.Build(Set)
	return nil
}

func NewScriptLoader(*grapicmd.Ctx) module.ScriptLoader {
	wire.Build(Set)
	return nil
}

func NewToolRepository(*grapicmd.Ctx) (tool.Repository, error) {
	wire.Build(Set)
	return nil, nil
}

func NewProtocWrapper(*grapicmd.Ctx) (protoc.Wrapper, error) {
	wire.Build(Set)
	return nil, nil
}

func NewInitializeProjectUsecase(*grapicmd.Ctx) usecase.InitializeProjectUsecase {
	wire.Build(Set)
	return nil
}
