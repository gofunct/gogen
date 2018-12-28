package template

import (
	"github.com/gofunct/gocookiecutter/gocookie"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate files from template directory",
	RunE: func(c *cobra.Command, args []string) error {
		var err error
		cookie, err = gocookie.NewGoCookieConfig()
		if err != nil {
			return err
		}
		return cookie.ExecTemplates()
	},
}



