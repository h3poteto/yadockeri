package services

import (
	"github.com/h3poteto/yadockeri/app/domains/branch"
	"github.com/h3poteto/yadockeri/app/domains/project"
	"github.com/h3poteto/yadockeri/lib/helm"
)

func Delete(project *project.Project, branch *branch.Branch) (string, error) {
	stackName := branch.GetStacName()
	deploy, err := helm.New(stackName, project.Namespace, false)
	if err != nil {
		return "", err
	}
	res, err := deploy.Delete(stackName)
	if err != nil {
		return "", err
	}
	return res.Info, nil
}
