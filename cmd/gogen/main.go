package main

import (
	"fmt"
	"github.com/gofunct/common/files"
	"github.com/gofunct/common/io"
	"github.com/gofunct/gogen/context"
	"github.com/gofunct/gogen/gogen/cmd"
	"github.com/gofunct/gogen/gogen"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmd.NewGogenCommand(&gogen.Ctx{
		IO:      io.Stdio(),
		RootDir: files.RootDir{files.Path(cwd)},
		Build: context.Build{
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
