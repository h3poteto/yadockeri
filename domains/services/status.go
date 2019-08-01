package services

import (
	"github.com/lapras-inc/yadockeri/domains/branch"
	"github.com/lapras-inc/yadockeri/lib/helm"
)

func ReleaseStatus(branch *branch.Branch) (string, error) {
	stackName := branch.GetStacName()

	deploy, err := helm.New(stackName, "", "")
	if err != nil {
		return "", err
	}
	release, err := deploy.Status()
	if err != nil {
		return "", err
	}
	return helm.PrintStatus(release)
}
