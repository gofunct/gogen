// +build tools

package main

// tool dependencies
import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/haya14busa/reviewdog/cmd/reviewdog"
	_ "github.com/izumin5210/gex/cmd/gex"
	_ "github.com/jessevdk/go-assets-builder"
	_ "github.com/kisielk/errcheck"
	_ "github.com/mitchellh/gox"
	_ "github.com/srvc/wraperr/cmd/wraperr"
	_ "golang.org/x/lint/golint"
	_ "honnef.co/go/tools/cmd/megacheck"
	_ "mvdan.cc/unparam"
)
