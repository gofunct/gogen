package context

import (
	"bytes"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"go/format"
	"path/filepath"
	"text/template"
)

type Entry struct {
	Template *template.Template
	Path     string
}

func (e *Entry) Create(fs afero.Fs, params interface{}) error {
	dir := filepath.Dir(e.Path)
	if ok, err := afero.DirExists(fs, dir); err != nil {
		return err
	} else if !ok {
		zap.L().Debug("create a directory", zap.String("dir", dir))
		err = fs.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	buf := new(bytes.Buffer)
	err := e.Template.Execute(buf, params)
	if err != nil {
		return err
	}

	data := buf.Bytes()

	if filepath.Ext(e.Path) == ".go" {
		data, err = format.Source(data)
		if err != nil {
			return err
		}
	}

	zap.L().Debug("create a new file", zap.String("path", e.Path))
	err = afero.WriteFile(fs, e.Path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
