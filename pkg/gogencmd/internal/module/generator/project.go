package generator

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module/generator/template"
	"github.com/gofunct/gogen/pkg/gogencmd/util/fs"
)

type projectGenerator struct {
	baseGenerator
}

func newProjectGenerator(fs afero.Fs, ui cli.UI) module.ProjectGenerator {
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
