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
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"io"
	"log"
	"os"
)

const layout = "20018-01-02"

func init() {
	gitCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "date [format: yyyy-mm-dd]")
}

var gitCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a github repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Filesystem abstraction based on memory
		fs := memfs.New()
		// Git objects storer based on memory
		storer := memory.NewStorage()
		// Clones the repository into the worktree (fs) and storer all the .git
		// content into the storer
		_, err := git.Clone(storer, fs, &git.CloneOptions{
			URL: url,
		})
		if err != nil {
			log.Fatal(err)
		}

		// Prints the content of the CHANGELOG file from the cloned repository
		changelog, err := fs.Open("CHANGELOG")
		if err != nil {
			log.Fatal(err)
		}

		io.Copy(os.Stdout, changelog)
	},
}
