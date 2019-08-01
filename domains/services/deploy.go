package services

import (
	"github.com/imdario/mergo"
	"github.com/lapras-inc/yadockeri/domains/branch"
	"github.com/lapras-inc/yadockeri/domains/project"
	"github.com/lapras-inc/yadockeri/domains/user"
	"github.com/lapras-inc/yadockeri/lib/helm"
	"github.com/sirupsen/logrus"
)

func DeployBranch(user *user.User, project *project.Project, branch *branch.Branch, revision string) (string, error) {
	stackName := branch.GetStacName()
	logrus.Infof("Deploy target stack: %s", stackName)

	filepath, err := helm.GitClone(project.HelmRepositoryUrl, user.Identifier, user.OauthToken)
	if err != nil {
		return "", err
	}
	logrus.Infof("Chart is downloaded: %s", filepath)

	deploy, err := helm.New(stackName, "", "")
	if err != nil {
		return "", err
	}
	version, err := deploy.Version()
	if err != nil {
		return "", err
	}
	logrus.Infof("helm version: %s", version)

	overrides := map[string]interface{}{}
	for _, v := range project.ValueOptions {
		mergo.Merge(&overrides, v.ToMap())
	}

	release, err := deploy.Status()
	if err != nil {
		// Install new chart if there is no release.
		res, err := deploy.NewRelease(filepath+"/"+project.HelmDirectoryName, project.Namespace, revision, overrides)
		if err != nil {
			return "", err
		}
		return deploy.PrintRelease(res)
	}
	// Update the release if there is already exist.
	res, err := deploy.UpdateRelease(release.Name, filepath+"/"+project.HelmDirectoryName, revision, overrides)
	if err != nil {
		return "", err
	}
	return deploy.PrintRelease(res)
}