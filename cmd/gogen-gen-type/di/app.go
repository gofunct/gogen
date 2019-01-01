package di

import (
	"github.com/gofunct/gogen/pkg/gencmd"
	"github.com/gofunct/gogen/pkg/protoc"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	Protoc protoc.Wrapper
}
