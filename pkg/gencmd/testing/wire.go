//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/gencmd"
)

func NewTestApp(*gencmd.Command, cli.UI) (*gencmd.App, error) {
	wire.Build(
		gencmd.Set,
	)
	return nil, nil
}
