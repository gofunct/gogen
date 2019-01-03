package generator

import (
	"github.com/gofunct/common/gogencmd/util/fs"
	"github.com/gofunct/gogen/module"
	"github.com/gofunct/gogen/module/generator/template"
	""github.com/gofunct/common/ui""
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type projectGenerator struct {
	baseGenerator
}

func newProjectGenerator(fs afero.Fs, ui ui.Menu) module.ProjectGenerator {
	return &projectGenerator{
		baseGenerator: newBaseGenerator(template.Init, fs, ui),
	}
}

func (g *projectGenerator) GenerateProject(rootDir, pkgName string) error {
	importPath, err := fs.GetImportPath(rootDir)
	if err != nil {
		return errors.WithStack(err)
	}

	if pkgName == "" {
		pkgName, err = fs.GetPackageName(rootDir)
		if err != nil {
			return errors.Wrap(err, "failed to decide a package name")
		}
	}

	data := map[string]interface{}{
		"packageName": pkgName,
		"importPath":  importPath,
	}
	return g.Generate(rootDir, data, generationConfig{})
}
