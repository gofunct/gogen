//+build wireinject

package cobrafs

import (
	"github.com/gofunct/common/ui"
	"github.com/google/wire"
)

func newApp(*Command) (*App, error) {
	wire.Build(
		CobraFsSet,
		ui.NewUI,
	)
	return nil, nil
}
