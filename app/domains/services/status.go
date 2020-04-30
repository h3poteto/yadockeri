package services

import (
	"github.com/h3poteto/yadockeri/app/domains/branch"
	"github.com/h3poteto/yadockeri/lib/helm"
)

func ReleaseStatus(branch *branch.Branch) (string, error) {
	stackName := branch.GetStacName()

	deploy, err := helm.New(stackName, false)
	if err != nil {
		return "", err
	}
	release, err := deploy.Status()
	if err != nil {
		return "", err
	}
	return helm.PrintStatus(release)
}
