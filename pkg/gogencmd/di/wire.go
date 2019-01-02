//+build wireinject

package di

import (
	"github.com/gofunct/common/bingen/tool"
	"github.com/google/wire"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/excmd"
	"github.com/gofunct/gogen/pkg/gogencmd"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/usecase"
	"github.com/gofunct/gogen/pkg/protoc"
)

func NewUI(*gogencmd.Ctx) cli.UI {
	wire.Build(Set)
	return nil
}

func NewCommandExecutor(*gogencmd.Ctx) excmd.Executor {
	wire.Build(Set)
	return nil
}

func NewGenerator(*gogencmd.Ctx) module.Generator {
	wire.Build(Set)
	return nil
}

func NewScriptLoader(*gogencmd.Ctx) module.ScriptLoader {
	wire.Build(Set)
	return nil
}

func NewToolRepository(*gogencmd.Ctx) (tool.Repository, error) {
	wire.Build(Set)
	return nil, nil
}

func NewProtocWrapper(*gogencmd.Ctx) (protoc.Wrapper, error) {
	wire.Build(Set)
	return nil, nil
}

func NewInitializeProjectUsecase(*gogencmd.Ctx) usecase.InitializeProjectUsecase {
	wire.Build(Set)
	return nil
}
