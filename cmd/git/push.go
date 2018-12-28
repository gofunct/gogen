// Copyright © 2018 Coleman Word <coleman.word@gofunct.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os/exec"
)

func init() {
	GitCmd.AddCommand(pushCmd)
}

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		{
			c := exec.Command("git", "add", ".")
			stderr, err := c.StderrPipe()
			if err != nil {
				log.Fatalf("%s %s", "failed to stage files", err)
			}

			err = c.Start()
			if err != nil {
				log.Fatalf("%s, %s", "failed to pipe errors", err)
			}
			logger.UI.Running("Waiting for command to finish...")
			out, _ := ioutil.ReadAll(stderr)
			logger.UI.Output(fmt.Sprintf("%s\n", out))

			err = c.Wait()
			if err != nil {
				log.Fatalf("%s %s", "failed to run command", err)
			}

			logger.UI.Success("Files added succesfully!")
		}

		{
			c := exec.Command("git", "commit", "-m", commitMsg)
			stderr, err := c.StderrPipe()
			if err != nil {
				log.Fatalf("%s, %s", "failed to pipe errors", err)
			}

			err = c.Start()
			if err != nil {
				log.Fatal("failed to commit files", err)
			}
			logger.UI.Running("Waiting for command to finish...")
			out, _ := ioutil.ReadAll(stderr)
			logger.UI.Output(fmt.Sprintf("%s\n", out))

			err = c.Wait()
			if err != nil {
				log.Fatalf("%s, %s", "failed to run command", err)
			}
			logger.UI.Success("Files commited succesfully!")
		}

		{
			c := exec.Command("git", "push", "origin", "master")
			stderr, err := c.StderrPipe()
			if err != nil {
				log.Fatalf("%s, %s", "failed to pipe errors", err)
			}

			if err != nil {
				log.Fatal("failed to push commit", err)
			}
			logger.UI.Running("Waiting for command to finish...")
			out, _ := ioutil.ReadAll(stderr)
			logger.UI.Output(fmt.Sprintf("%s\n", out))

			err = c.Wait()
			if err != nil {
				log.Fatalf("%s, %s", "failed to run command", err)
			}
			logger.UI.Success("Files pushed finished succesfully!")
		}

	},
}
