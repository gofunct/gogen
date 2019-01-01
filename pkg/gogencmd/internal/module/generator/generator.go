package generator

import (
	"github.com/spf13/afero"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module"
)

// New creates a module.Generator instance.
func New(fs afero.Fs, ui cli.UI) module.Generator {
	return &generator{
		ProjectGenerator: newProjectGenerator(fs, ui),
	}
}

type generator struct {
	module.ProjectGenerator
}
