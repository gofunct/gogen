package template

import (
	"github.com/gofunct/gocookiecutter/gocookie"
	"github.com/spf13/cobra"
)


// templateCmd represents the template command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate files from template directory",
	Run: func(c *cobra.Command, args []string) {
		gocookie.Cookie.ExecTemplates(tmplPath)
	},
}



