package di

import (
	"github.com/gofunct/common/cli"
	"github.com/gofunct/common/excmd"
	"github.com/gofunct/common/gogencmd/module/script"
	"github.com/gofunct/gogen/gogen"
	"github.com/gofunct/gogen/module"
	"github.com/gofunct/gogen/module/generator"
	"github.com/gofunct/gogen/protoc"
	"github.com/gofunct/gogen/usecase"
	"github.com/google/wire"
	"github.com/izumin5210/gex"
)

func ProvideGenerator(ctx *gogen.Ctx, ui cli.UI) module.Generator {
	return generator.New(
		ctx.FS,
		ui,
	)
}

func ProvideScriptLoader(ctx *gogen.Ctx, executor excmd.Executor) module.ScriptLoader {
	return script.NewLoader(ctx.FS, executor, ctx.RootDir.String())
}

func ProvideInitializeProjectUsecase(ctx *gogen.Ctx, gexCfg *gex.Config, ui cli.UI, generator module.Generator) usecase.InitializeProjectUsecase {
	return usecase.NewInitializeProjectUsecase(
		ui,
		generator,
		gexCfg,
	)
}

var Set = wire.NewSet(
	gogen.CtxSet,
	protoc.WrapperSet,
	cli.UIInstance,
	excmd.NewExecutor,
	ProvideGenerator,
	ProvideScriptLoader,
	ProvideInitializeProjectUsecase,
)
