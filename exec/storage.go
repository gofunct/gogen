package exec

import (
	cntext "context"
	"github.com/gofunct/common/cloud/storage"
	"github.com/gofunct/gogen/context"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

// NewVersionCommand create a new cobra.Command to print the version information.
func NewUploadCommand(cfg context.Build) *cobra.Command {
	return &cobra.Command{
		Use:           "upload",
		Short:         "Print the version information",
		Long:          "Print the version information",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, _ []string) {
			var (
				file string
			)

			// Define our input.
			cmd.Flags().StringVarP(&cfg.Cloud, "cloud", "c", "", "Cloud storage to use")
			cmd.Flags().StringVarP(&cfg.StorageBucket, "bucket", "b", "go-cloud-bucket", "Name of bucket")
			cmd.Flags().StringVarP(&file, "file", "f", "", "file target")

			ctx := cntext.Background()
			// Open a connection to the bucket.
			b, err := storage.SetupBucket(cntext.Background(), cfg.Cloud, cfg.StorageBucket)
			if err != nil {
				log.Fatalf("Failed to setup bucket: %s", err)
			}

			// Prepare the file for upload.
			data, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalf("Failed to read file: %s", err)
			}

			w, err := b.NewWriter(ctx, file, nil)
			if err != nil {
				log.Fatalf("Failed to obtain writer: %s", err)
			}
			_, err = w.Write(data)
			if err != nil {
				log.Fatalf("Failed to write to bucket: %s", err)
			}
			if err = w.Close(); err != nil {
				log.Fatalf("Failed to close: %s", err)
			}

		},
	}
}

// NewVersionCommand create a new cobra.Command to print the version information.
func NewDownloadCommand(cfg context.Build) *cobra.Command {
	return &cobra.Command{
		Use:           "download",
		Short:         "Print the version information",
		Long:          "Print the version information",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, _ []string) {
			var (
				key string
			)

			// Define our input.
			cmd.Flags().StringVarP(&cfg.Cloud, "cloud", "c", "", "Cloud storage to use")
			cmd.Flags().StringVarP(&cfg.StorageBucket, "bucket", "b", "go-cloud-bucket", "Name of bucket")
			cmd.Flags().StringVarP(&key, "key", "f", "", "file target")

			ctx := cntext.Background()
			// Open a connection to the bucket.
			bucket, err := storage.SetupBucket(cntext.Background(), cfg.Cloud, cfg.StorageBucket)
			if err != nil {
				log.Fatalf("Failed to setup bucket: %s", err)
			}

			if err = storage.CopyFileFromBucket(bucket, ctx, key); err != nil {
				log.Fatalf("Failed to setup bucket: %s", err)
			}
		},
	}
}
