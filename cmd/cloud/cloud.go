package cloud

import "github.com/spf13/cobra"

var (
	cloudProvider string
)

func init() {
	CloudCmd.PersistentFlags().StringVarP(&cloudProvider, "cloud", "c", "", "Cloud storage to use")
}

var CloudCmd = &cobra.Command{
	Use: "cloud opts",
	Short: "options for interacting with aws and google cloud",
}
