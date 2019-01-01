package gencmd

import (
	"github.com/gofunct/gogen/pkg/gogencmd"
	"github.com/pkg/errors"
)

func defaultCtx() *Ctx {
	return &Ctx{
		Ctx: &gogencmd.Ctx{},
	}
}

// Ctx defines a context of a generator.
type Ctx struct {
	*gogencmd.Ctx

	CreateAppFunc CreateAppFunc
}

func (c *Ctx) apply(opts []Option) {
	for _, f := range opts {
		f(c)
	}
}

// CreateApp initializes dependencies.
func (c *Ctx) CreateApp(cmd *Command) (*App, error) {
	f := c.CreateAppFunc
	if c.CreateAppFunc == nil {
		f = newApp
	}
	app, err := f(cmd)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return app, nil
}
