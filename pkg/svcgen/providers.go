package svcgen

import (
	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/gogencmd"
	"github.com/google/wire"
	"github.com/gofunct/gogen/pkg/protoc"
	"github.com/gofunct/gogen/pkg/svcgen/params"

)

func ProvideParamsBuilder(rootDir cli.RootDir, protocCfg *protoc.Config, gogenCfg *gogencmd.Config) params.Builder {
	return params.NewBuilder(
		rootDir,
		protocCfg.ProtosDir,
		protocCfg.OutDir,
		gogenCfg.Gogen.ServerDir,
		gogenCfg.Package,
	)
}

var Set = wire.NewSet(
	ProvideParamsBuilder,
	App{},
)
