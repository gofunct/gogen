package cli

import (
	"github.com/gofunct/common/files"
)

// RootDir represents a project root directory.
type RootDir struct {
	files.Path
}

// BinDir returns the directory path contains executable binaries.
func (d *RootDir) BinDir() files.Path {
	return d.Join("bin")
}
