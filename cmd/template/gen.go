package template

import (
	"github.com/gofunct/gogen/config"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate files from template directory",
	RunE: func(c *cobra.Command, args []string) error {
		var err error
		cookie, err = config.NewGoCookieConfig()
		if err != nil {
			return err
		}
		return cookie.ExecTemplates()
	},
}



