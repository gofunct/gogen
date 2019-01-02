package build

import (
	"github.com/gofunct/common/files"
	"github.com/gofunct/common/io"
	"github.com/gofunct/common/logging"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"k8s.io/utils/exec"
)

func NewDefaultCliCommand(wd files.Path, build Build) *cobra.Command {
	return NewCliCommand(&Ctx{
		WorkingDir: wd,
		IO:         io.Stdio(),
		FS:         afero.NewOsFs(),
		Exec:       exec.New(),
		Build:      build,
	})
}

func NewCliCommand(ctx *Ctx) *cobra.Command {
	cmd := &cobra.Command{
		Use: ctx.Build.AppName,
		Short: ctx.Build.Description,
	}

	logging.AddLoggingFlags(cmd)

	cmd.AddCommand(
		newInitCommand(ctx),
		NewVersionCommand(ctx.IO, ctx.Build),
	)

	return cmd
}
