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
	"github.com/gofunct/gogen/cmd/cloud"
	ccobra "github.com/gofunct/gogen/cmd/cobra"
	"github.com/gofunct/gogen/cmd/docker"
	"github.com/gofunct/gogen/cmd/git"
	"github.com/gofunct/gogen/cmd/protoc"
	"github.com/gofunct/gogen/cmd/template"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var (
	logger, _ = zap.NewDevelopment()
	goPath = os.Getenv("GOPATH")
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gogen",
	Short: "A dev utitility tool for golang based projects",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logging.AddLoggingFlags(RootCmd)
	{
		RootCmd.AddCommand(git.GitCmd)
		RootCmd.AddCommand(docker.DockerCmd)
		RootCmd.AddCommand(protoc.ProtocCmd)
		RootCmd.AddCommand(template.TemplateCmd)
		RootCmd.AddCommand(ccobra.CobraCmd)
		RootCmd.AddCommand(cloud.CloudCmd)
	}

	{
		RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	}
}
