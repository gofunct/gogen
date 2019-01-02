package main

import (
	"context"
	"github.com/gofunct/common/files"
	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/gencmd"
	"github.com/gofunct/gogen/pkg/gogencmd"
	"github.com/gofunct/gogen/pkg/svcgen"
	"github.com/gofunct/gogen/pkg/protoc"
	"testing"

	"github.com/spf13/afero"

	gencmdtesting "github.com/gofunct/gogen/pkg/gencmd/testing"
	svcgentesting "github.com/gofunct/gogen/pkg/svcgen/testing"
)

func TestRun(t *testing.T) {
	cases := []svcgentesting.Case{
		{
			Test:  "simple",
			GArgs: []string{"foo"},
			DArgs: []string{"foo"},
			Files: []string{
				"api/protos/foo.proto",
				"app/server/foo_server.go",
				"app/server/foo_server_register_funcs.go",
				"app/server/foo_server_test.go",
			},
		},
		{
			Test:  "specify package",
			GArgs: []string{"foo"},
			DArgs: []string{"foo"},
			Files: []string{
				"api/protos/foo.proto",
				"app/server/foo_server.go",
				"app/server/foo_server_register_funcs.go",
				"app/server/foo_server_test.go",
			},
			PkgName: "testcompany.testapp",
		},
		{
			Test:  "nested",
			GArgs: []string{"foo/bar"},
			DArgs: []string{"foo/bar"},
			Files: []string{
				"api/protos/foo/bar.proto",
				"app/server/foo/bar_server.go",
				"app/server/foo/bar_server_register_funcs.go",
				"app/server/foo/bar_server_test.go",
			},
		},
		{
			Test:  "nested with specify pacakge",
			GArgs: []string{"foo/bar"},
			DArgs: []string{"foo/bar"},
			Files: []string{
				"api/protos/foo/bar.proto",
				"app/server/foo/bar_server.go",
				"app/server/foo/bar_server_register_funcs.go",
				"app/server/foo/bar_server_test.go",
			},
			PkgName: "testcompany.testapp",
		},
		{
			Test:  "snake_case name",
			GArgs: []string{"foo/bar_baz"},
			DArgs: []string{"foo/bar_baz"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test:  "kebab-case name",
			GArgs: []string{"foo/bar-baz"},
			DArgs: []string{"foo/bar-baz"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test:  "with some standard methods",
			GArgs: []string{"foo/bar-baz", "list", "create", "delete"},
			DArgs: []string{"foo/bar-baz"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test:  "with non-standard methods",
			GArgs: []string{"foo/bar-baz", "list", "create", "rename", "delete", "move_move"},
			DArgs: []string{"foo/bar-baz"},
			Files: []string{
				"api/protos/foo/bar_baz.proto",
				"app/server/foo/bar_baz_server.go",
				"app/server/foo/bar_baz_server_register_funcs.go",
				"app/server/foo/bar_baz_server_test.go",
			},
		},
		{
			Test:  "specify proto dir",
			GArgs: []string{"qux"},
			DArgs: []string{"qux"},
			Files: []string{
				"pkg/foo/protos/qux.proto",
				"app/server/qux_server.go",
				"app/server/qux_server_register_funcs.go",
				"app/server/qux_server_test.go",
			},
			ProtoDir: "pkg/foo/protos",
		},
		{
			Test:  "specify proto out dir",
			GArgs: []string{"quux"},
			DArgs: []string{"quux"},
			Files: []string{
				"api/protos/quux.proto",
				"app/server/quux_server.go",
				"app/server/quux_server_register_funcs.go",
				"app/server/quux_server_test.go",
			},
			ProtoOutDir: "api/out",
		},
		{
			Test:  "specify server dir",
			GArgs: []string{"corge"},
			DArgs: []string{"corge"},
			Files: []string{
				"api/protos/corge.proto",
				"pkg/foo/server/corge_server.go",
				"pkg/foo/server/corge_server_register_funcs.go",
				"pkg/foo/server/corge_server_test.go",
			},
			ServerDir: "pkg/foo/server",
		},
		{
			Test:  "skip tests",
			GArgs: []string{"--skip-test", "book"},
			DArgs: []string{"book"},
			Files: []string{
				"api/protos/book.proto",
				"app/server/book_server.go",
				"app/server/book_server_register_funcs.go",
			},
			SkippedFiles: map[string]struct{}{
				"app/server/book_server_test.go": {},
			},
		},
		{
			Test:  "specify resource name",
			GArgs: []string{"library", "--resource-name=book"},
			DArgs: []string{"library"},
			Files: []string{
				"api/protos/library.proto",
				"app/server/library_server.go",
				"app/server/library_server_register_funcs.go",
				"app/server/library_server_test.go",
			},
		},
	}

	rootDir := cli.RootDir{files.Path("/home/src/testapp")}

	createSvcApp := func(cmd *gencmd.Command) (*svcgen.App, error) {
		return svcgentesting.NewTestApp(cmd, &fakeProtocWrapper{}, cli.NopUI)
	}
	createGenApp := func(cmd *gencmd.Command) (*gencmd.App, error) {
		return gencmdtesting.NewTestApp(cmd, cli.NopUI)
	}
	createCmd := func(t *testing.T, fs afero.Fs, tc svcgentesting.Case) gencmd.Executor {
		ctx := &gogencmd.Ctx{
			FS:      fs,
			RootDir: rootDir,
			Config: gogencmd.Config{
				Package: tc.PkgName,
			},
			ProtocConfig: protoc.Config{
				ProtosDir: tc.ProtoDir,
				OutDir:    tc.ProtoOutDir,
			},
		}
		ctx.Config.Gogen.ServerDir = tc.ServerDir
		return buildCommand(createSvcApp, gencmd.WithGrapiCtx(ctx), gencmd.WithCreateAppFunc(createGenApp))
	}

	ctx := &svcgentesting.Ctx{
		GOPATH:    "/home",
		RootDir:   rootDir,
		CreateCmd: createCmd,
		Cases:     cases,
	}

	svcgentesting.Run(t, ctx)
}

type fakeProtocWrapper struct{}

func (*fakeProtocWrapper) Exec(context.Context) error { return nil }
