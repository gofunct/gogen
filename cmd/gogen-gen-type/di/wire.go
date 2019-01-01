//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/gencmd"
	"github.com/gofunct/gogen/pkg/protoc"
)

func NewApp(*gencmd.Command) (*App, error) {
	wire.Build(
		App{},
		gencmd.Set,
		cli.UIInstance,
		protoc.WrapperSet,
	)
	return nil, nil
}
