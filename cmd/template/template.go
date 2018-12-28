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

package template

import (
	"github.com/gofunct/common/logging"
	"github.com/prometheus/common/log"
	"github.com/gofunct/gocookiecutter/gocookie"
	"github.com/spf13/cobra"
	"os"
)

var (
	logger = logging.NewLogger(os.Stdout)
	goPath = os.Getenv("GOPATH")
	destPath string
	cookie *gocookie.GoCookieConfig
)

func init() {
	var err error
	cookie, err = gocookie.NewGoCookieConfig()
	if err != nil {
		log.Fatal("failed to inititalize gocookie config", err)
	}
	TemplateCmd.AddCommand(genCmd)
	TemplateCmd.AddCommand(funcCmd)
	TemplateCmd.PersistentFlags().StringVarP(&destPath, "dest-path", "d", ".", "specify the path to the output directory")
}

// templateCmd represents the template command
var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Template opts",
}

