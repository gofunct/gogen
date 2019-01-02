package template

import (
	"github.com/gofunct/gogen/pkg/config"
	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var funcCmd = &cobra.Command{
	Use:   "func",
	Short: "list sprig template functions",
	Run: func(c *cobra.Command, args []string) {
		config.List()
	},
}


