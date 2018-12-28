package template

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate files from template directory",
	RunE: func(c *cobra.Command, args []string) error {
		return cookie.ExecTemplates()
	},
}



