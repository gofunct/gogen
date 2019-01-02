package main

import (
	"fmt"
	"github.com/gofunct/common/files"
	"github.com/gofunct/common/io"
	"github.com/gofunct/gogen/pkg/gogencmd"
	"os"


	"github.com/gofunct/gogen/pkg/cli"
	"github.com/gofunct/gogen/pkg/gogencmd/cmd"
	"github.com/gofunct/common/build"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmd.NewGogenCommand(&gogencmd.Ctx{
		IO:      io.Stdio(),
		RootDir: cli.RootDir{files.Path(cwd)},
		Build: build.Build{
			AppName:   name,
			Version:   version,
			Revision:  revision,
			BuildDate: buildDate,
		},
	}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
