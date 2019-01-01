package cmd

import (
	"github.com/gofunct/common/build"
	"github.com/gofunct/common/logging"
	"github.com/gofunct/gogen/pkg/gogencmd"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/gofunct/common/io"
)

// NewGogenCommand creates a new command object.
func NewGogenCommand(ctx *gogencmd.Ctx) *cobra.Command {
	initErr := ctx.Init()

	cmd := &cobra.Command{
		Use:           ctx.Build.AppName,
		Short:         "JSON API framework implemented with gRPC and Gateway",
		Long:          "",
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
		newInitCommand(ctx),
		newProtocCommand(ctx),
		newBuildCommand(ctx),
		build.NewVersionCommand(ctx.IO, ctx.Build),
	)
	cmd.AddCommand(newGenerateCommands(ctx)...)
	cmd.AddCommand(newUserDefinedCommands(ctx)...)

	return cmd
}
