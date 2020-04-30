package services

import (
	"github.com/h3poteto/yadockeri/app/domains/branch"
	"github.com/h3poteto/yadockeri/app/domains/project"
	"github.com/h3poteto/yadockeri/app/domains/user"
	"github.com/h3poteto/yadockeri/app/values"
	"github.com/h3poteto/yadockeri/lib/helm"
	"github.com/sirupsen/logrus"
)

// DeployBranch deploy a branch using project's helm chart.
// At first, determine stack name from branch name.
// Then git clone helm chart repository.
// Generate override variables and create helm release.
func DeployBranch(user *user.User, project *project.Project, branch *branch.Branch, variable *values.TemplateVariable) (string, error) {
	stackName := branch.GetStacName()
	logrus.Infof("Deploy target stack: %s", stackName)

	filepath, err := helm.GitClone(project.HelmRepositoryURL, user.Identifier, user.OauthToken)
	if err != nil {
		return "", err
	}
	logrus.Infof("Chart is downloaded: %s", filepath)

	deploy, err := helm.New(stackName, false)
	if err != nil {
		return "", err
	}

	overrides := []string{}
	for _, v := range project.ValueOptions {
		overrides = append(overrides, v.ToString())
	}

	overrides, err = variable.ReplaceVariablesAll(overrides)
	if err != nil {
		return "", err
	}

	release, err := deploy.Status()
	if err != nil {
		// Install new chart if there is no release.
		res, err := deploy.NewRelease(filepath+"/"+project.HelmDirectoryName, project.Namespace, overrides)
		if err != nil {
			return "", err
		}
		return deploy.PrintRelease(res)
	}
	// Update the release if there is already exist.
	res, err := deploy.UpdateRelease(release.Name, filepath+"/"+project.HelmDirectoryName, overrides)
	if err != nil {
		return "", err
	}
	return deploy.PrintRelease(res)
}
