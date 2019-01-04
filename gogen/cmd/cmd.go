package cmd

import (
	"github.com/gofunct/common/errors"
	"github.com/gofunct/common/io"
	"github.com/gofunct/common/logging"
	"github.com/gofunct/gogen/exec"
	"github.com/gofunct/gogen/gogen"
	"github.com/spf13/cobra"
)

func NewGogenCommand(ctx *gogen.Ctx) *cobra.Command {
	initErr := ctx.Init()

	cmd := &cobra.Command{
		Use:           ctx.Build.AppName,
		Short:         ctx.Build.Description,
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return errors.WithStack(initErr)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			io.Close()
		},
	}

	logging.AddLoggingFlags(cmd)

	cmd.AddCommand(
		NewInitCommand(ctx),
		NewProtocCommand(ctx),
		NewBuildCommand(ctx),
		exec.NewVersionCommand(ctx.IO, ctx.Build),
	)
	cmd.AddCommand(newGenerateCommands(ctx)...)
	cmd.AddCommand(NewUserDefinedCommands(ctx)...)

	return cmd
}
