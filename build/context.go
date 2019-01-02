package build

import (
	"github.com/gofunct/common/files"
	"github.com/gofunct/common/io"
	"github.com/spf13/afero"
	"k8s.io/utils/exec"
)

type Ctx struct {
	WorkingDir files.Path
	IO         io.IO
	FS         afero.Fs
	Exec       exec.Interface

	Build Build
}
