package services

import (
	"github.com/h3poteto/yadockeri/app/domains/branch"
	"github.com/h3poteto/yadockeri/lib/helm"
)

func Delete(branch *branch.Branch) (string, error) {
	stackName := branch.GetStacName()
	deploy, err := helm.New(stackName, false)
	if err != nil {
		return "", err
	}
	res, err := deploy.Delete(stackName)
	if err != nil {
		return "", err
	}
	return res.Info, nil
}
