package cobrafs

import (
	"github.com/gofunct/gogen/gogen"
)

// Option configures a command context.
type Option func(*Ctx)

// WithGrapiCtx specifies a grapi command context.
func WithGrapiCtx(gctx *gogen.Ctx) Option {
	return func(ctx *Ctx) {
		ctx.Ctx = gctx
	}
}

// WithCreateAppFunc specifies a dependencies initializer.
func WithCreateAppFunc(f CreateAppFunc) Option {
	return func(ctx *Ctx) {
		ctx.CreateAppFunc = f
	}
}
