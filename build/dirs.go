package build

import (
	"github.com/gofunct/common/files"
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
