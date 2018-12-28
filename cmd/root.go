// Copyright © 2018 Coleman Word <coleman.word@gofunct.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/gofunct/common/logging"
	"github.com/gofunct/gocookiecutter/cmd/docker"
	"github.com/gofunct/gocookiecutter/cmd/git"
	"github.com/gofunct/gocookiecutter/cmd/protoc"
	"github.com/gofunct/gocookiecutter/cmd/template"
	"github.com/spf13/cobra"
	"log"
	"os"
	kitlog "github.com/go-kit/kit/log"
)

var (
	logger = logging.NewLogger()
	goPath = os.Getenv("GOPATH")
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gocookiecutter",
	Short: "A brief description of your application",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logger.AddColor()
	log.SetOutput(kitlog.NewStdlibAdapter(logger.KitLog))
	{
		RootCmd.AddCommand(git.GitCmd)
		RootCmd.AddCommand(docker.DockerCmd)
		RootCmd.AddCommand(protoc.ProtocCmd)
		RootCmd.AddCommand(template.TemplateCmd)
	}

	{
		RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	}
}
