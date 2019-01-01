package cobra

import (
	ccmd "github.com/gofunct/common/cli/cmd"
	"github.com/gofunct/common/files"
	"github.com/spf13/cobra"
	"os"
	"github.com/gofunct/common/build"
)

var genCmd = &cobra.Command{
	Use: "gen",
	Short: "generate a new cobra based application",
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},

}


func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	c := ccmd.NewDefaultCligCommand(files.Path(wd), build.Build{
		AppName:   appName,
		Version:   version,
		Revision:  revision,
		BuildDate: buildDate,
	})

	return c.Execute()
}

