//+build wireinject

package gocloud

import (
	"context"
	"github.com/gofunct/gogen/gocloud/google"

	"github.com/google/wire"
	"gocloud.dev/blob"
	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/gcp"
	"gocloud.dev/gcp/gcpcloud"
	"gocloud.dev/mysql/cloudmysql"
	"gocloud.dev/runtimevar"
	pb "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
)

// This file wires the generic interfaces up to Google Cloud Platform (GCP). It
// won't be directly included in the final binary, since it includes a Wire
// injector template function (setupGCP), but the declarations will be copied
// into wire_gen.go when Wire is run.

// setupGCP is a Wire injector function that sets up the application using GCP.
func SetupGCP(ctx context.Context, flags *cliFlags) (*Application, func(), error) {
	// This will be filled in by Wire with providers from the provider sets in
	// wire.Build.
	wire.Build(
		gcpcloud.GCP,
		cloudmysql.Open,
		ApplicationSet,
		GcpBucket,
		GcpRuntimeConfig,
		GcpSQLParams,
	)
	return nil, nil, nil
}

// gcpBucket is a Wire provider function that returns the GCS bucket based on
// the command-line flags.
func GcpBucket(ctx context.Context, flags *cliFlags, client *gcp.HTTPClient) (*blob.Bucket, error) {
	return gcsblob.OpenBucket(ctx, flags.bucket, client, nil)
}

// gcpSQLParams is a Wire provider function that returns the Cloud SQL
// connection parameters based on the command-line flags. Other providers inside
// gcpcloud.GCP use the parameters to construct a *sql.DB.
func GcpSQLParams(id gcp.ProjectID, flags *cliFlags) *cloudmysql.Params {
	return &cloudmysql.Params{
		ProjectID: string(id),
		Region:    flags.cloudSQLRegion,
		Instance:  flags.dbHost,
		Database:  flags.dbName,
		User:      flags.dbUser,
		Password:  flags.dbPassword,
	}
}

// gcpMOTDVar is a Wire provider function that returns the Message of the Day
// variable from Runtime Configurator.
func GcpRuntimeConfig(ctx context.Context, client pb.RuntimeConfigManagerClient, project gcp.ProjectID, flags *cliFlags) (*runtimevar.Variable, func(), error) {
	name := google.ResourceName{
		ProjectID: string(project),
		Config:    flags.runtimeConfigName,
		Variable:  flags.runVar,
	}
	v, err := google.NewVariable(client, name, runtimevar.StringDecoder, &google.Options{
		WaitDuration: flags.runVarWaitTime,
	})
	if err != nil {
		return nil, nil, err
	}
	return v, func() { v.Close() }, nil
}
