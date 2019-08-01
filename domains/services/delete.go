package services

import (
	"github.com/lapras-inc/yadockeri/domains/branch"
	"github.com/lapras-inc/yadockeri/lib/helm"
)

func Delete(branch *branch.Branch) (string, error) {
	stackName := branch.GetStacName()
	deploy, err := helm.New(stackName, "", "")
	if err != nil {
		return "", err
	}
	res, err := deploy.Delete(stackName)
	if err != nil {
		return "", err
	}
	return res.Info, nil
}
