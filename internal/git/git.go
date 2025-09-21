package git

import (
	"dotwaifu/internal/config"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
)

func InitRepository() error {
	configDir := config.GetConfigDir()
	_, err := git.PlainInit(configDir, false)
	return err
}

func IsGitRepository() bool {
	configDir := config.GetConfigDir()
	_, err := git.PlainOpen(configDir)
	return err == nil
}

func AddAndCommit(message string) error {
	configDir := config.GetConfigDir()
	repo, err := git.PlainOpen(configDir)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = worktree.AddGlob(".")
	if err != nil {
		return err
	}

	_, err = worktree.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "dotwaifu",
			Email: "dotwaifu@local",
			When:  time.Now(),
		},
	})

	return err
}

func GetStatus() (git.Status, error) {
	configDir := config.GetConfigDir()
	repo, err := git.PlainOpen(configDir)
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	return worktree.Status()
}