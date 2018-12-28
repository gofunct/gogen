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
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
	"syscall"
	"time"
)

const layout = "20018-01-02"

func init() {
	now := time.Now()
	gitCmd.PersistentFlags().StringVarP(
		&date, "date", "d", now.Format(layout), "date [format: yyyy-mm-dd]")
	gitCmd.PersistentFlags().StringVarP(
		&repoDir, "repo", "r", "",
		"local directory path for clone GitHub repository (required)")
	gitCmd.PersistentFlags().StringVarP(
		&commitMsg, "commit", "c", ":herb: ʕ ◔ϖ◔ʔ :herb:", "commit message")
	gitCmd.PersistentFlags().StringVarP(
		&name, "name", "n", "ʕ ◔ϖ◔ʔ", "commit author name")
	gitCmd.PersistentFlags().StringVarP(
		&mail, "mail", "m", "kusa@example.com", "commit author mail address")
}


var (
	repoDir   string
	date      string
	commitMsg string
	name      string
	mail      string
	gitCmd   = &cobra.Command{
		Use:   "create",
		Short: "Create GitHub contribution",
		Long:  "Create GitHub contribution on date specified by date option",
		Run: func(cmd *cobra.Command, args []string) {
			if err := goGit(); err != nil {
				logger.Error("create GitHub contribution error", zap.Error(err))
			}
		},
	}
)

func goGit() error {
	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		logger.Error("open git repository error", zap.String("dir_path", repoDir), zap.Error(err))
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		logger.Error("failed to get work tree", zap.Error(err))
		return err
	}

	time, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		logger.Error("time parse error", zap.String("date", date), zap.Error(err))
		return err
	}

	logger.Info("execute commit",
		zap.String("name", name), zap.String("e-mail", mail),
		zap.String("date", date), zap.String("commit_message", commitMsg))
	commit, err := wt.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  name,
			Email: mail,
			When:  time,
		},
	})
	if err != nil {
		logger.Error("commit error", zap.Error(err))
		return err
	}
	logger.Info("complete commit", zap.String("commit_hash", commit.String()))

	logger.Info("input credential")
	auth, err := inputCredentials()
	if err != nil {
		logger.Error("failed read credentials", zap.Error(err))
		// rollback reset empty commit
		rollbackCommit(wt, commit)
		return err
	}

	logger.Info("execute push", zap.String("repository", repoDir))
	if err := repo.Push(&git.PushOptions{
		Auth:     auth,
		Progress: os.Stdout,
	}); err != nil {
		logger.Error("push error", zap.Error(err))
		// rollback reset empty commit
		rollbackCommit(wt, commit)
		return err
	}
	logger.Info("complete push")

	return nil
}

func inputCredentials() (transport.AuthMethod, error) {
	fmt.Print("user name: ")
	user, err := terminal.ReadPassword(syscall.Stdin)
	// new line
	fmt.Println()
	if err != nil {
		logger.Error("failed read user name", zap.Error(err))
		return nil, err
	}

	fmt.Print("password: ")
	pass, err := terminal.ReadPassword(syscall.Stdin)
	// new line
	fmt.Println()
	if err != nil {
		logger.Error("failed read password", zap.Error(err))
		return nil, err
	}

	return &http.BasicAuth{
		Username: string(user),
		Password: string(pass),
	}, nil
}

func rollbackCommit(wt *git.Worktree, commit plumbing.Hash) {
	logger.Info("rollback reset commit", zap.String("commit_hash", commit.String()))
	if err := wt.Reset(&git.ResetOptions{
		Commit: commit,
	}); err != nil {
		logger.Error("failed rollback reset commit",
			zap.String("commit_hash", commit.String()), zap.Error(err))
	}
	logger.Info("complete reset commit", zap.String("commit_hash", commit.String()))
}