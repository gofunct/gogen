package di

import (
	"github.com/gofunct/gogen/pkg/gogencmd"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module/script"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/usecase"
	"github.com/gofunct/gogen/pkg/protoc"
	"github.com/google/wire"
	"github.com/gofunct/common/bingen"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/excmd"
)

func ProvideGenerator(ctx *gogencmd.Ctx, ui cli.UI) module.Generator {
	return generator.New(
		ctx.FS,
		ui,
	)
}

func ProvideScriptLoader(ctx *gogencmd.Ctx, executor excmd.Executor) module.ScriptLoader {
	return script.NewLoader(ctx.FS, executor, ctx.RootDir.String())
}

func ProvideInitializeProjectUsecase(ctx *gogencmd.Ctx, gexCfg *bingen.Config, ui cli.UI, generator module.Generator) usecase.InitializeProjectUsecase {
	return usecase.NewInitializeProjectUsecase(
		ui,
		generator,
		gexCfg,
	)
}

var Set = wire.NewSet(
	gogencmd.CtxSet,
	protoc.WrapperSet,
	cli.UIInstance,
	excmd.NewExecutor,
	ProvideGenerator,
	ProvideScriptLoader,
	ProvideInitializeProjectUsecase,
)
