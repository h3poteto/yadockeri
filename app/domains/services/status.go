package services

import (
	"github.com/h3poteto/yadockeri/app/domains/branch"
	"github.com/h3poteto/yadockeri/app/domains/project"
	"github.com/h3poteto/yadockeri/lib/helm"
)

// ReleaseStatus gets a status of released helm package related project and branch.
func ReleaseStatus(project *project.Project, branch *branch.Branch) (string, error) {
	stackName := branch.GetStacName()

	deploy, err := helm.New(stackName, project.Namespace, false)
	if err != nil {
		return "", err
	}
	release, err := deploy.Status()
	if err != nil {
		return "", err
	}
	return helm.PrintStatus(release)
}
