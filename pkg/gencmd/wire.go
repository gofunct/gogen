//+build wireinject

package gencmd

import (
	"github.com/google/wire"
	"github.com/gofunct/gogen/pkg/cli"
)

func newApp(*Command) (*App, error) {
	wire.Build(
		Set,
		cli.UIInstance,
	)
	return nil, nil
}
