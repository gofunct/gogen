package generator

import (
	"github.com/gofunct/gogen/module"
	""github.com/gofunct/common/ui""
	"github.com/spf13/afero"
)

// New creates a module.Generator instance.
func New(fs afero.Fs, ui ui.Menu) module.Generator {
	return &generator{
		ProjectGenerator: newProjectGenerator(fs, ui),
	}
}

type generator struct {
	module.ProjectGenerator
}
