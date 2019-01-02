package context

import (
	"github.com/gofunct/common/errors"
	"github.com/gofunct/common/files"
	"path/filepath"
	"strings"
)

type RootDir struct {
	files.Path
}

func (d *RootDir) BinDir() files.Path {
	return d.Join("bin")
}

func (d *RootDir) CmdDir() files.Path {
	return d.Join("cmd")
}

func (d *RootDir) TmplDir() files.Path {
	return d.Join("templates")
}

func (d *RootDir) ProtoDir() files.Path {
	return d.Join("proto")
}

func (d *RootDir) VendorDir() files.Path {
	return d.Join("vendor")
}

func getImportPath(rootPath string) (importPath string, err error) {
	for _, gopath := range filepath.SplitList(BuildContext.GOPATH) {
		prefix := filepath.Join(gopath, "src") + string(filepath.Separator)
		// FIXME: should not use strings.HasPrefix
		if strings.HasPrefix(rootPath, prefix) {
			importPath = filepath.ToSlash(strings.Replace(rootPath, prefix, "", 1))
			break
		}
	}
	if importPath == "" {
		err = errors.New("failed to get the import path")
	}
	return
}
