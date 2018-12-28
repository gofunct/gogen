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

package docker

import (
	"github.com/gofunct/common/logging"
	"github.com/spf13/cobra"
	"os"
)

var (
	dockerEndpoint string
	logger = logging.NewLogger()
	goPath = os.Getenv("GOPATH")
)

func init() {
	DockerCmd.PersistentFlags().StringVarP(&dockerEndpoint, "endpoint", "e", "unix:///var/run/docker.sock", "docker endpoint")
}

// dockerCmd represents the up command
var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker opts",
	}
