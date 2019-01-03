package inject

import (
	"github.com/gofunct/common/executor"
	"github.com/gofunct/common/ui"
	"github.com/gofunct/gogen/gogen"
	"github.com/gofunct/gogen/module"
	"github.com/gofunct/gogen/module/generator"
	"github.com/gofunct/gogen/module/script"
	"github.com/gofunct/gogen/protoc"
	"github.com/gofunct/gogen/usecase"
	"github.com/google/wire"
	"github.com/izumin5210/gex"
)

func ProvideGenerator(ctx *gogen.Ctx, ui ui.UI) module.Generator {
	return generator.New(
		ctx.FS,
		ui,
	)
}

func ProvideScriptLoader(ctx *gogen.Ctx, executor executor.Executor) module.ScriptLoader {
	return script.NewLoader(ctx.FS, executor, ctx.RootDir.String())
}

func ProvideInitializeProjectUsecase(ctx *gogen.Ctx, gexCfg *gex.Config, ui ui.UI, generator module.Generator) usecase.InitializeProjectUsecase {
	return usecase.NewInitializeProjectUsecase(
		ui,
		generator,
		gexCfg,
	)
}

var InjectorSet = wire.NewSet(
	gogen.CtxSet,
	protoc.WrapperSet,
	executor.NewExecutor,
	ProvideGenerator,
	ProvideScriptLoader,
	ProvideInitializeProjectUsecase,
	ui.NewUI,
)
