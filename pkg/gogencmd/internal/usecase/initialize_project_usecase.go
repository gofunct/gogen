package usecase

import (
	"context"

	"github.com/gofunct/common/bingen"
	"github.com/pkg/errors"

	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/gogencmd/internal/module"
)

// InitializeProjectUsecase is an interface to create a new grapi project.
type InitializeProjectUsecase interface {
	Perform(rootDir string, cfg InitConfig) error
	GenerateProject(rootDir, pkgName string) error
	InstallDeps(rootDir string, cfg InitConfig) error
}

// NewInitializeProjectUsecase creates a new InitializeProjectUsecase instance.
func NewInitializeProjectUsecase(ui cli.UI, generator module.ProjectGenerator, binCfg *bingen.Config) InitializeProjectUsecase {
	return &initializeProjectUsecase{
		ui:        ui,
		generator: generator,
		binCfg:    binCfg,
	}
}

type initializeProjectUsecase struct {
	ui        cli.UI
	generator module.ProjectGenerator
	binCfg    *bingen.Config
}

func (u *initializeProjectUsecase) Perform(rootDir string, cfg InitConfig) error {
	u.ui.Section("Initialize project")

	var err error
	err = u.GenerateProject(rootDir, cfg.Package)
	if err != nil {
		return errors.Wrap(err, "failed to initialize project")
	}

	u.ui.Subsection("Install dependencies")
	err = u.InstallDeps(rootDir, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to execute `dep ensure`")
	}

	return nil
}

func (u *initializeProjectUsecase) GenerateProject(rootDir, pkgName string) error {
	return errors.WithStack(u.generator.GenerateProject(rootDir, pkgName))
}

func (u *initializeProjectUsecase) InstallDeps(rootDir string, cfg InitConfig) error {
	u.binCfg.WorkingDir = rootDir
	repo, err := u.binCfg.Create()
	if err == nil {
		spec := cfg.BuildSpec()
		err = repo.Add(
			context.TODO(),
			"github.com/izumin5210/grapi/cmd/grapi"+spec,
			"github.com/izumin5210/grapi/cmd/grapi-gen-command",
			"github.com/izumin5210/grapi/cmd/grapi-gen-service",
			"github.com/izumin5210/grapi/cmd/grapi-gen-scaffold-service",
			"github.com/izumin5210/grapi/cmd/grapi-gen-type",
			// TODO: make configurable
			"github.com/golang/protobuf/protoc-gen-go",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger",
		)
	}
	return errors.WithStack(err)
}
