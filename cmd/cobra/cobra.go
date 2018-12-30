package cobra

import (
	"github.com/gofunct/common/clib"
	ccmd "github.com/gofunct/common/clig/cmd"
	"github.com/spf13/cobra"
	"os"
)

const (
	appName = "cobra"
	version = "v0.2.0"
)

var (
	revision, buildDate string
)


var CobraCmd = &cobra.Command{
	Use: "init",
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

	c := ccmd.NewDefaultCligCommand(clib.Path(wd), clib.Build{
		AppName:   appName,
		Version:   version,
		Revision:  revision,
		BuildDate: buildDate,
	})

	return c.Execute()
}
