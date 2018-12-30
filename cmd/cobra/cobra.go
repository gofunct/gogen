package cobra

import "github.com/spf13/cobra"

var (
	appName, version, revision, buildDate string
)

func init() {
	CobraCmd.AddCommand(genCmd)
	CobraCmd.PersistentFlags().StringVarP(&appName, "app", "a", "", "the desired cli app name")
	CobraCmd.PersistentFlags().StringVarP(&appName, "version", "v", "0.1.0", "the version of the cli application")
}

var CobraCmd = &cobra.Command{
	Use:  "cobra",
	Short: "cobra cli opts",
}