package cobrafs

import (
	"github.com/gofunct/common/errors"
	"github.com/gofunct/gogen/gogen"
)

func defaultCtx() *Ctx {
	return &Ctx{
		Ctx: &gogen.Ctx{},
	}
}

// Ctx defines a context of a generator.
type Ctx struct {
	*gogen.Ctx

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
