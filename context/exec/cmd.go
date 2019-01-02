package exec

import (
	"github.com/gofunct/common/files"
	"github.com/gofunct/common/io"
	"github.com/gofunct/common/logging"
	"github.com/gofunct/gogen/context"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"k8s.io/utils/exec"
)

func NewDefaultCliCommand(wd files.Path) *cobra.Command {
	cfg := context.NewConfig()
	return NewCliCommand(&context.Ctx{
		WorkingDir: wd,
		IO:         io.Stdio(),
		FS:         afero.NewOsFs(),
		Exec:       exec.New(),
		Build:      context.BuildFromConfig(cfg),
		Config: 	cfg,
	})
}

func NewCliCommand(ctx *context.Ctx) *cobra.Command {
	cmd := &cobra.Command{
		Use: ctx.Build.AppName,
		Short: ctx.Build.Description,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return ctx.Config.ReadConfig()
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return ctx.Config.ReadConfig()
		},
	}

	logging.AddLoggingFlags(cmd)

	cmd.AddCommand(
		context.NewInitCommand(ctx),
		NewVersionCommand(ctx.IO, ctx.Build),

		//TODO: Add command: gcloud storage upload
		//TODO: Add command: gcloud storage download
		//TODO: Add command: gcloud kubernetes manifest deployment
		//TODO: Add command: gcloud cloudsql cli
		//TODO: Add command: gcloud cloudsql cli
	)

	return cmd
}
