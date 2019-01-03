package cobrafs

import (
	"github.com/gofunct/gogen/gogen"
	"net/http"

	"github.com/google/wire"
)

func ProvideGogenCtx(ctx *Ctx) *gogen.Ctx            { return ctx.Ctx }
func ProvideCtx(cmd *Command) *Ctx                   { return cmd.Ctx() }
func ProvideTemplateFS(cmd *Command) http.FileSystem { return cmd.TemplateFS }
func ProvideShouldRun(cmd *Command) ShouldRunFunc    { return cmd.ShouldRun }

// Set contains providers for DI.
var CobraFsSet = wire.NewSet(
	gogen.CtxSet,
	ProvideGogenCtx,
	ProvideCtx,
	ProvideTemplateFS,
	ProvideShouldRun,
	NewGenerator,
	App{},
)
