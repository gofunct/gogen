package cmd

import (
	"github.com/gofunct/common/errors"
	"github.com/gofunct/gogen/gogen"
	"github.com/gofunct/gogen/gogen/inject"
	"github.com/gofunct/gogen/usecase"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"path/filepath"
)

func NewInitCommand(ctx *gogen.Ctx) *cobra.Command {
	var cfg usecase.InitConfig

	cmd := &cobra.Command{
		Use:           "init [name]",
		Short:         "Initialize a gogen application",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root, err := parseInitArgs(ctx, args)
			if err != nil {
				return errors.WithStack(err)
			}
			zap.L().Debug("parseInitArgs", zap.String("root", root))

			return errors.WithStack(inject.NewInitializeProjectUsecase(ctx).Perform(root, cfg))
		},
	}

	cmd.PersistentFlags().StringVar(&cfg.Revision, "revision", "", "Specify gogen revision")
	cmd.PersistentFlags().StringVar(&cfg.Branch, "branch", "", "Specify gogen branch")
	cmd.PersistentFlags().StringVar(&cfg.Version, "version", ctx.Build.Version, "Specify gogen version")
	cmd.PersistentFlags().BoolVar(&cfg.HEAD, "HEAD", false, "Use HEAD gogen")
	cmd.PersistentFlags().StringVarP(&cfg.Package, "package", "p", "", `Package name of the application(default: "<parent_package_or_username>.<app_name>")`)

	return cmd
}

func parseInitArgs(ctx *gogen.Ctx, args []string) (root string, err error) {
	if argCnt := len(args); argCnt != 1 {
		err = errors.Errorf("invalid argument count: want 1, got %d", argCnt)
		return
	}

	arg := args[0]
	root = ctx.RootDir.String()

	if arg == "." {
		return
	}
	root = arg
	if !filepath.IsAbs(arg) {
		root = ctx.RootDir.Join(arg).String()
	}
	return
}
