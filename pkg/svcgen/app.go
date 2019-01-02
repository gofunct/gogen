package svcgen

import (
	"github.com/gofunct/gogen/pkg/gencmd"
	"github.com/gofunct/gogen/pkg/protoc"
	"github.com/gofunct/gogen/pkg/svcgen/params"

)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	ProtocWrapper protoc.Wrapper
	ParamsBuilder params.Builder
}
