package project

import "github.com/lapras-inc/yadockeri/values"

type Project struct {
	ID                int
	UserID            int
	Title             string
	BaseURL           string
	RepositoryOwner   string
	RepositoryName    string
	HelmRepositoryUrl string
	HelmDirectoryName string
	Namespace         string
	ValueOptions      []*values.OverrideValue
}

func (p *Project) CheckOwner(userID int) bool {
	return p.UserID == userID
}
