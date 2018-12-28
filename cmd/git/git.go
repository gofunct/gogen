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

package git

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"github.com/gofunct/common/logging"
	kitlog "github.com/go-kit/kit/log"
)

var (
	goPath = os.Getenv("GOPATH")
	remoteUrl			string
	commitMsg 			string
	logger = logging.NewLogger()
)

func init() {
	log.SetOutput(kitlog.NewStdlibAdapter(logger.KitLog))
	GitCmd.PersistentFlags().StringVarP(&remoteUrl, "url", "u", "", "remote url of target repo")
	GitCmd.PersistentFlags().StringVar(&commitMsg, "msg", "default msg",  "remote url of target repo")
}

var GitCmd = &cobra.Command{
	Use:   "git",
	Short: "git opts",
}