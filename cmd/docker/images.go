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

package docker

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"

	"github.com/spf13/cobra"
)

func init() {
	DockerCmd.AddCommand(imagesCmd)
}

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "list docker images",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := docker.NewClient(dockerEndpoint)
		if err != nil {
			panic(err)
		}
		imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
		if err != nil {
			panic(err)
		}
		for _, img := range imgs {
			fmt.Println("ID: ", img.ID)
			fmt.Println("RepoTags: ", img.RepoTags)
			fmt.Println("Created: ", img.Created)
			fmt.Println("Size: ", img.Size)
			fmt.Println("VirtualSize: ", img.VirtualSize)
			fmt.Println("ParentId: ", img.ParentID)
		}
	},
}

