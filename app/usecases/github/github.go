package github

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func client(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)
}

func GetRepos(token string) ([]*github.Repository, error) {
	cli := client(token)
	nextPage := -1
	var repositories []*github.Repository
	for nextPage != 0 {
		if nextPage < 0 {
			nextPage = 0
		}
		repositoryOption := &github.RepositoryListOptions{
			Type:      "all",
			Sort:      "full_name",
			Direction: "asc",
			ListOptions: github.ListOptions{
				Page:    nextPage,
				PerPage: 50,
			},
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		repos, res, err := cli.Repositories.List(ctx, "", repositoryOption)
		nextPage = res.NextPage
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, repos...)
	}
	return repositories, nil
}

func GetBranches(owner, repo, token string) ([]*github.Branch, error) {
	cli := client(token)
	nextPage := -1
	var branches []*github.Branch
	for nextPage != 0 {
		if nextPage < 0 {
			nextPage = 0
		}
		branchOption := &github.ListOptions{
			Page:    nextPage,
			PerPage: 50,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		b, res, err := cli.Repositories.ListBranches(ctx, owner, repo, branchOption)
		nextPage = res.NextPage
		if err != nil {
			return nil, err
		}
		branches = append(branches, b...)
	}
	return branches, nil
}
