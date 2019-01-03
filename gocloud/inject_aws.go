//+build wireinject

package gocloud

import (
	"context"
	awsclient "github.com/aws/aws-sdk-go/aws/client"
	"github.com/gofunct/gogen/gocloud/aws"
	"github.com/google/wire"
	"gocloud.dev/aws/awscloud"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
	"gocloud.dev/mysql/rdsmysql"
	"gocloud.dev/runtimevar"
)

// This file wires the generic interfaces up to Amazon Web Services (AWS). It
// won't be directly included in the final binary, since it includes a Wire
// injector template function (setupAWS), but the declarations will be copied
// into wire_gen.go when Wire is run.

// setupAWS is a Wire injector function that sets up the application using AWS.
func SetupAWS(ctx context.Context, flags *cliFlags) (*Application, func(), error) {
	// This will be filled in by Wire with providers from the provider sets in
	// wire.Build.
	wire.Build(
		awscloud.AWS,
		rdsmysql.Open,
		ApplicationSet,
		AwsBucket,
		AwsRuntimeConfig,
		AwsSQLParams,
	)
	return nil, nil, nil
}

// awsBucket is a Wire provider function that returns the S3 bucket based on the
// command-line flags.
func AwsBucket(ctx context.Context, cp awsclient.ConfigProvider, flags *cliFlags) (*blob.Bucket, error) {
	return s3blob.OpenBucket(ctx, flags.bucket, cp, nil)
}

// awsSQLParams is a Wire provider function that returns the RDS SQL connection
// parameters based on the command-line flags. Other providers inside
// awscloud.AWS use the parameters to construct a *sql.DB.
func AwsSQLParams(flags *cliFlags) *rdsmysql.Params {
	return &rdsmysql.Params{
		Endpoint: flags.dbHost,
		Database: flags.dbName,
		User:     flags.dbUser,
		Password: flags.dbPassword,
	}
}

// awsMOTDVar is a Wire provider function that returns the Message of the Day
// variable from SSM Parameter Store.
func AwsRuntimeConfig(ctx context.Context, sess awsclient.ConfigProvider, flags *cliFlags) (*runtimevar.Variable, error) {
	return aws.NewVariable(sess, flags.runVar, runtimevar.StringDecoder, &aws.Options{
		WaitDuration: flags.runVarWaitTime,
	})
}
