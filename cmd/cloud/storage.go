package cloud

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"context"
	"github.com/gofunct/common/cloud/storage"
)

var (
	bucket string
	file string
)

func init() {
	storageCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "Name of bucket")
	storageCmd.Flags().StringVarP(&file, "file", "f", "", "Name of file to upload")
}

var storageCmd = &cobra.Command{
	Use: "storage",
	Short: "cloud storage options",
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()
		// Open a connection to the bucket.
		b, err := storage.SetupBucket(context.Background(), cloudProvider, bucket)
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
