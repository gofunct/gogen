package context

import (
	"github.com/spf13/cobra"
	"go/build"
)

var (
	BuildContext = build.Default
)

func NewInitCommand(c *Ctx) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "init",
		Args: cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			if err := c.ParseEntries(); err != nil {
				return err
			}
			return c.RunDepAndGex()
		},
	}

	cmd.Flags().String("name", c.Build.AppName, "name of cli application")
	return cmd
}
