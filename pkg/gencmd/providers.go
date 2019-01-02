package gencmd

import (
	"github.com/gofunct/gogen/pkg/gogencmd"
	"net/http"

	"github.com/google/wire"

)

func ProvideGogenCtx(ctx *Ctx) *gogencmd.Ctx         { return ctx.Ctx }
func ProvideCtx(cmd *Command) *Ctx                   { return cmd.Ctx() }
func ProvideTemplateFS(cmd *Command) http.FileSystem { return cmd.TemplateFS }
func ProvideShouldRun(cmd *Command) ShouldRunFunc    { return cmd.ShouldRun }

// Set contains providers for DI.
var Set = wire.NewSet(
	gogencmd.CtxSet,
	ProvideGogenCtx,
	ProvideCtx,
	ProvideTemplateFS,
	ProvideShouldRun,
	NewGenerator,
	App{},
)
