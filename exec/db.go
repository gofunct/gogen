package exec

import (
	"flag"
	"github.com/gofunct/common/gocloud/db"
	"github.com/gofunct/common/io"
	"github.com/gofunct/gogen/context"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {

}

// NewVersionCommand create a new cobra.Command to print the version information.
func NewDBCommand(io io.IO, cfg context.Build) *cobra.Command {
	return &cobra.Command{
		Use:           "version",
		Short:         "Print the version information",
		Long:          "Print the version information",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, _ []string) {
			log.SetFlags(0)
			log.SetPrefix("gcp/provision_db: ")
			project := flag.String("project", "", "GCP project ID")
			serviceAccount := flag.String("service_account", "", "name of service account in GCP project")
			instance := flag.String("instance", "", "database instance name")
			database := flag.String("database", "", "name of database to initialize")
			password := flag.String("password", "", "root password for the database")
			schema := flag.String("schema", "", "path to .sql file defining the database schema")
			flag.Parse()
			missing := false
			flag.VisitAll(func(f *flag.Flag) {
				if f.Value.String() == "" {
					log.Printf("Required flag -%s is not set.", f.Name)
					missing = true
				}
			})
			if missing {
				os.Exit(64)
			}
			if err := db.ProvisionDB(*project, *serviceAccount, *instance, *database, *password, *schema); err != nil {
				log.Fatal(err)
			}
		},
	}
}
