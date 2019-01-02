package protoc

import (
	"github.com/gofunct/common/bingen"
	"github.com/gofunct/common/bingen/tool"
	"github.com/gofunct/common/io"
	"github.com/gofunct/common/logging"
	"sync"

	"github.com/google/wire"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"k8s.io/utils/exec"

	"github.com/gofunct/gogen/pkg/cli"
)

var (
	binCfg   *bingen.Config
	binCfgMu sync.Mutex

	toolRepo   tool.Repository
	toolRepoMu sync.Mutex
)

func ProvideBingenConfig(
	fs afero.Fs,
	execer exec.Interface,
	io io.IO,
	rootDir cli.RootDir,
) *bingen.Config {
	binCfgMu.Lock()
	defer binCfgMu.Unlock()
	if binCfg == nil {
		binCfg = &bingen.Config{
			OutWriter:  io.Out(),
			ErrWriter:  io.Err(),
			InReader:   io.In(),
			FS:         fs,
			Execer:     execer,
			WorkingDir: rootDir.String(),
			Verbose:    logging.IsVerbose() || logging.IsDebug(),
			Logger:     zap.NewStdLog(zap.L()),
		}
	}
	return binCfg
}

func ProvideToolRepository(binCfg *bingen.Config) (tool.Repository, error) {
	toolRepoMu.Lock()
	defer toolRepoMu.Unlock()
	if toolRepo == nil {
		var err error
		toolRepo, err = binCfg.Create()
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return toolRepo, nil
}

// WrapperSet is a provider set that includes gex things and Wrapper instance.
var WrapperSet = wire.NewSet(
	ProvideBingenConfig,
	ProvideToolRepository,
	NewWrapper,
)
