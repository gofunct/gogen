package script

import (
	"context"
	"path/filepath"

	"github.com/gofunct/common/executor"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

type script struct {
	fs            afero.Fs
	executor      executor.Executor
	rootDir       string
	name, binPath string
	srcPaths      []string
}

func (s *script) Name() string {
	return s.name
}

func (s *script) Build(args ...string) error {
	zap.L().Debug("build script", zap.String("name", s.name), zap.String("bin", s.binPath), zap.Strings("srcs", s.srcPaths))
	err := fs.CreateDirIfNotExists(s.fs, filepath.Dir(s.binPath))
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = s.executor.Exec(context.TODO(), "go", s.buildOpts(args)...)
	if err != nil {
		return errors.Wrapf(err, "failed to build %v", s.srcPaths)
	}

	return nil
}

func (s *script) Run(args ...string) error {
	_, err := s.executor.Exec(
		context.TODO(),
		s.binPath,
		executor.WithArgs(args...),
		executor.WithDir(s.rootDir),
		executor.WithIOConnected(),
	)
	return errors.WithStack(err)
}

func (s *script) buildOpts(args []string) []executor.Option {
	built := make([]string, 0, 3+len(args)+len(s.srcPaths))
	built = append(built, "build", "-o="+s.binPath)
	built = append(built, args...)
	built = append(built, s.srcPaths...)
	return []executor.Option{
		executor.WithArgs(built...),
		executor.WithDir(s.rootDir),
		executor.WithIOConnected(),
	}
}
