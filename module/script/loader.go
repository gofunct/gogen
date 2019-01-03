package script

import (
	"github.com/gofunct/common/errors"
	"github.com/gofunct/common/executor"
	"github.com/gofunct/common/files"
	"github.com/gofunct/gogen/module"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"path/filepath"
	"runtime"
	"sort"
)

// NewLoader creates a new ScriptLoader instance.
func NewLoader(fs afero.Fs, excmd executor.Executor, rootDir string) module.ScriptLoader {
	return &scriptLoader{
		fs:      fs,
		excmd:   excmd,
		rootDir: rootDir,
		binDir:  filepath.Join(rootDir, "bin"),
		scripts: make(map[string]module.Script),
	}
}

type scriptLoader struct {
	fs      afero.Fs
	excmd   executor.Executor
	rootDir string
	binDir  string
	scripts map[string]module.Script
	names   []string
}

func (f *scriptLoader) Load(dir string) error {
	srcsByDir, err := files.FindMainPackagesAndSources(f.fs, dir)
	zap.L().Debug("found main packages", zap.Any("srcs_by_dir", srcsByDir))
	if err != nil {
		return errors.Wrap(err, "failed to find commands")
	}
	for dir, srcs := range srcsByDir {
		srcPaths := make([]string, 0, len(srcs))
		for _, name := range srcs {
			srcPaths = append(srcPaths, filepath.Join(dir, name))
		}
		name := filepath.Base(dir)
		ext := ""
		if runtime.GOOS == "windows" {
			ext = ".exe"
		}
		f.scripts[name] = &script{
			fs:       f.fs,
			executor: nil,
			rootDir:  f.rootDir,
			name:     name,
			binPath:  filepath.Join(f.binDir, name+ext),
			srcPaths: srcPaths,
		}
		f.names = append(f.names, name)
	}
	sort.Strings(f.names)
	return nil
}

func (f *scriptLoader) Get(name string) (script module.Script, ok bool) {
	script, ok = f.scripts[name]
	return
}

func (f *scriptLoader) Names() []string {
	return f.names
}
