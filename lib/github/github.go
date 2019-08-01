package github

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHub struct {
	client *github.Client
}

func New(token string) *GitHub {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return &GitHub{client: client}
}

func (g *GitHub) GetRevision(owner, repository, branch string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	gitBranch, _, err := g.client.Repositories.GetBranch(ctx, owner, repository, branch)
	if err != nil {
		return "", err
	}
	return *gitBranch.Commit.SHA, nil
}
