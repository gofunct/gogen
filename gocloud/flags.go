package gocloud

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"time"
)

type cliFlags struct {
	bucket            string
	dbHost            string
	dbName            string
	dbUser            string
	dbPassword        string
	runVar           string
	runVarWaitTime   time.Duration
	listen            string
	cloudSQLRegion    string
	runtimeConfigName string
}

var envFlag string

func NewGoCloudCommand(config *viper.Viper) *cobra.Command {
	cf  := new(cliFlags)

	cmd := &cobra.Command{
		Use:   "gocloud",
		Short: "cloud opts",
	}

	cmd.PersistentFlags().DurationVar(&cf.runVarWaitTime, "runtime_var_wait", 5*time.Second, "polling frequency of message of the day")
	config.SetDefault("runtime_var_wait", 5*time.Second)

	{
		cmd.PersistentFlags().StringVar(&envFlag, "env", "local", "what is environment do you want to run under?(gcp or aws)")
		config.SetDefault("env", envFlag)
	}
	{
		cmd.PersistentFlags().StringVar(&cf.listen, "listen", ":8080", "what port do you want to listen on?")
		config.SetDefault("listen", cf.listen)

	}
	{
		cmd.PersistentFlags().StringVar(&cf.bucket, "bucket", "", "what bucket name do you want to setup?")
		config.SetDefault("bucket", cf.bucket)

	}
	{
		cmd.PersistentFlags().StringVar(&cf.dbHost, "db_host", "", "what is the database host or Cloud SQL instance name?")
		config.SetDefault("db_host", cf.dbHost)

	}
	{
		cmd.PersistentFlags().StringVar(&cf.dbName, "db_name", "", "what is the database name?")
		config.SetDefault("db_name", cf.dbName)
	}
	{
		cmd.PersistentFlags().StringVar(&cf.dbUser, "db_user", "guestbook", "what is the database username?")
		config.SetDefault("db_user", cf.dbUser)

	}
	{
		cmd.PersistentFlags().StringVar(&cf.dbPassword, "db_password", "", "what is the database user password?")
		config.SetDefault("db_password", cf.dbPassword)

	}
	{
		cmd.PersistentFlags().StringVar(&cf.runVar, "runtime_var", "", "what is the runtime variable location?")
		config.SetDefault("runtime_var", cf.runVar)

	}
	{
		cmd.PersistentFlags().StringVar(&cf.cloudSQLRegion, "cloud_sql_region", "", "region of the Cloud SQL instance (GCP only)")
		config.SetDefault("cloud_sql_region", cf.cloudSQLRegion)
	}
	{
		cmd.PersistentFlags().StringVar(&cf.runtimeConfigName, "runtime_config", "", "runtime Configurator config resource (GCP only)")
		config.SetDefault("runtime_config", cf.runtimeConfigName)

	}

	config.BindPFlags(cmd.PersistentFlags())

	init := &cobra.Command{
		Use:   "init",
		Short: "initialize",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			var app *Application
			var cleanup func()
			var err error
			switch envFlag {
			case "gcp":
				app, cleanup, err = SetupGCP(ctx, cf)
			case "aws":
				if cf.dbPassword == "" {
					cf.dbPassword = "xyzzy"
				}
				app, cleanup, err = SetupAWS(ctx, cf)
			case "local":
				if cf.dbHost == "" {
					cf.dbHost = "localhost"
				}
				if cf.dbPassword == "" {
					cf.dbPassword = "xyzzy"
				}
				app, cleanup, err = SetupLocal(ctx, cf)
			default:
				log.Fatalf("unknown -env=%s", envFlag)
			}
			if err != nil {
				log.Fatal(err)
			}
			defer cleanup()

			// Set up URL routes.
			r := mux.NewRouter()
			//r.HandleFunc("/", app.index)
			//r.HandleFunc("/sign", app.sign)
			r.HandleFunc("/blob/{key:.+}", app.ServeBlob)

			// Listen and serve HTTP.
			log.Printf("Running, connected to %q cloud", envFlag)
			log.Fatal(app.srv.ListenAndServe(cf.listen, r))
		},
	}

	cmd.AddCommand(init)

	return cmd
}
