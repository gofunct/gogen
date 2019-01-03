package cmd

import (
	"github.com/gofunct/common/errors"
	"github.com/gofunct/common/ui"
	"github.com/gofunct/gogen/gogen"
	"github.com/gofunct/gogen/gogen/inject"
	"github.com/gofunct/gogen/module"
	"github.com/spf13/cobra"
)

func NewUserDefinedCommands(ctx *gogen.Ctx) (cmds []*cobra.Command) {
	if !ctx.IsInsideApp() {
		return
	}

	scriptLoader := inject.NewScriptLoader(ctx)

	err := scriptLoader.Load(ctx.RootDir.Join("cmd").String())
	if err != nil {
		// TODO: log
		return
	}

	ui := inject.NewUI(ctx)

	for _, name := range scriptLoader.Names() {
		cmds = append(cmds, NewUserDefinedCommand(ui, scriptLoader, name))
	}

	return
}

func NewUserDefinedCommand(ui ui.UI, scriptLoader module.ScriptLoader, name string) *cobra.Command {
	return &cobra.Command{
		Use:           name + " [-- BUILD_OPTIONS] [-- RUN_ARGS]",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(c *cobra.Command, args []string) (err error) {
			script, ok := scriptLoader.Get(name)
			if !ok {
				err = errors.Wrapf(err, "faild to find subcommand %s", name)
				return
			}

			pos := len(args)
			for i, arg := range args {
				if arg == "--" {
					pos = i
					break
				}
			}
			var buildArgs, runArgs []string
			if pos == len(args) {
				buildArgs = args
			} else {
				buildArgs = args[:pos]
				runArgs = args[pos+1:]
			}

			ui.Info("Building...")
			err = errors.WithStack(script.Build(buildArgs...))
			if err != nil {
				return
			}

			ui.Info("Starting...")
			return errors.WithStack(script.Run(runArgs...))
		},
	}
}
