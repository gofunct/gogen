package generator

import (
	"github.com/gofunct/common/ui"
	"github.com/gofunct/gogen/module"
	"github.com/spf13/afero"
)

// New creates a module.Generator instance.
func New(fs afero.Fs, ui ui.UI) module.Generator {
	return &generator{
		ProjectGenerator: newProjectGenerator(fs, ui),
	}
}

type generator struct {
	module.ProjectGenerator
}
