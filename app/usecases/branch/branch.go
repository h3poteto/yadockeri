package branch

import (
	"errors"

	"github.com/h3poteto/yadockeri/app/domains/services"
	"github.com/h3poteto/yadockeri/app/repositories/branches"
	"github.com/h3poteto/yadockeri/app/repositories/project_values"
	"github.com/h3poteto/yadockeri/app/repositories/projects"
	"github.com/h3poteto/yadockeri/app/repositories/users"
	"github.com/h3poteto/yadockeri/db"
	"github.com/h3poteto/yadockeri/lib/github"
	"github.com/sirupsen/logrus"
)

type Branch struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	StackName string `json:"stack_name"`
	URL       string `json:"url"`
}

func GetBranchesByProject(projectID int) ([]*Branch, error) {
	repository := branches.New(db.SharedInstance().Connection)
	branchList, err := repository.GetByProject(projectID)
	if err != nil {
		return nil, err
	}
	var results []*Branch
	for _, branch := range branchList {
		b := &Branch{
			ID:        branch.ID,
			ProjectID: branch.ProjectID,
			UserID:    branch.UserID,
			Name:      branch.Name,
			URL:       branch.URL,
			StackName: branch.GetStacName(),
		}
		results = append(results, b)
	}
	return results, nil
}

func CreateBranch(projectID, userID int, name string) (*Branch, error) {
	repository := branches.New(db.SharedInstance().Connection)
	id, err := repository.Create(projectID, userID, name)
	if err != nil {
		return nil, err
	}
	b, err := repository.GetByID(id)
	if err != nil {
		return nil, err

	}
	return &Branch{
		ID:        b.ID,
		ProjectID: b.ProjectID,
		UserID:    b.UserID,
		Name:      b.Name,
		URL:       b.URL,
		StackName: b.GetStacName(),
	}, nil
}

func GetBranch(projectID, id int) (*Branch, error) {
	repository := branches.New(db.SharedInstance().Connection)
	branch, err := repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if !branch.CheckProject(projectID) {
		return nil, errors.New("branch not found")
	}

	return &Branch{
		ID:        branch.ID,
		ProjectID: branch.ProjectID,
		UserID:    branch.UserID,
		Name:      branch.Name,
		URL:       branch.URL,
		StackName: branch.GetStacName(),
	}, nil
}

func Deploy(projectID, branchID int) (string, error) {
	branchRepository := branches.New(db.SharedInstance().Connection)
	b, err := branchRepository.GetByID(branchID)
	if err != nil {
		return "", err
	}
	if !b.CheckProject(projectID) {
		return "", errors.New("branch not found")
	}

	userRepository := users.New(db.SharedInstance().Connection)
	user, err := userRepository.GetByID(b.UserID)
	if err != nil {
		return "", err
	}
	projectRepository := projects.New(db.SharedInstance().Connection)
	project, err := projectRepository.GetByID(b.ProjectID)
	if err != nil {
		return "", err
	}
	valuesRepository := project_values.New(db.SharedInstance().Connection)
	overrides, err := valuesRepository.GetByProject(project.ID)
	if err != nil {
		return "", err
	}
	project.ValueOptions = overrides

	hub := github.New(user.OauthToken)
	revision, err := hub.GetRevision(project.RepositoryOwner, project.RepositoryName, b.Name)
	if err != nil {
		return "", err
	}
	logrus.Infof("Deploy target SHA1: %s", revision)

	res, err := services.DeployBranch(user, project, b, revision)
	if err != nil {
		return "", err
	}
	logrus.Info(res)

	b.UpdateURL(project.BaseURL)
	err = branchRepository.Update(b)
	if err != nil {
		return "", err
	}
	return res, err
}

func GetStatus(projectID, branchID int) (string, error) {
	branchRepository := branches.New(db.SharedInstance().Connection)
	branch, err := branchRepository.GetByID(branchID)
	if err != nil {
		return "", err
	}
	if !branch.CheckProject(projectID) {
		return "", errors.New("branch not found")
	}

	return services.ReleaseStatus(branch)
}

func Delete(projectID, branchID int) error {
	branchRepository := branches.New(db.SharedInstance().Connection)
	branch, err := branchRepository.GetByID(branchID)
	if err != nil {
		return err
	}
	if !branch.CheckProject(projectID) {
		return errors.New("branch not found")
	}

	output, err := services.Delete(branch)
	logrus.Info(output)
	// Sometimes branch is not deployed in helm.
	// So ignore error from helm.
	if err != nil {
		logrus.Warn(err)
	}
	err = branchRepository.DeleteByID(branch.ID)
	if err != nil {
		return err
	}
	return nil
}
