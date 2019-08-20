package projects

import (
	"database/sql"

	"github.com/h3poteto/yadockeri/app/domains/project"
)

// Projects is repository struct.
type Projects struct {
	db *sql.DB
}

// New returns a new Projects.
func New(db *sql.DB) *Projects {
	return &Projects{
		db,
	}
}

// All returns all projects.
func (p *Projects) All() ([]*project.Project, error) {
	var projects []*project.Project
	rows, err := p.db.Query("SELECT id, user_id, title, base_url, repository_owner, repository_name, helm_repository_url, helm_directory_name, namespace FROM projects")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, userID int
		var title, baseURL, repositoryOwner, repositoryName, helmRepositoryURL, helmDirectoryName, namespace string
		err = rows.Scan(&id, &userID, &title, &baseURL, &repositoryOwner, &repositoryName, &helmRepositoryURL, &helmDirectoryName, &namespace)
		if err != nil {
			return nil, err
		}
		p := &project.Project{
			ID:                id,
			UserID:            userID,
			Title:             title,
			BaseURL:           baseURL,
			RepositoryOwner:   repositoryOwner,
			RepositoryName:    repositoryName,
			HelmRepositoryURL: helmRepositoryURL,
			HelmDirectoryName: helmDirectoryName,
			Namespace:         namespace,
		}
		projects = append(projects, p)
	}
	return projects, nil
}

// Create is create method for a project.
func (p *Projects) Create(tx *sql.Tx, userID int, title, baseURL, repositoryOwner, repositoryName, helmRepositoryURL, helmDirectoryName, namespace string) (int, error) {
	var id int
	err := tx.QueryRow("INSERT INTO projects (user_id, title, base_url, repository_owner, repository_name, helm_repository_url, helm_directory_name, namespace) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", userID, title, baseURL, repositoryOwner, repositoryName, helmRepositoryURL, helmDirectoryName, namespace).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetByID gets a project from specified ID.
func (p *Projects) GetByID(id int) (*project.Project, error) {
	var findID, userID int
	var title, baseURL, repositoryOwner, repositoryName, helmRepositoryURL, helmDirectoryName, namespace string
	err := p.db.QueryRow("SELECT id, user_id, title, base_url, repository_owner, repository_name, helm_repository_url, helm_directory_name, namespace FROM projects WHERE id = $1", id).Scan(&findID, &userID, &title, &baseURL, &repositoryOwner, &repositoryName, &helmRepositoryURL, &helmDirectoryName, &namespace)
	if err != nil {
		return nil, err
	}
	return &project.Project{
		ID:                findID,
		UserID:            userID,
		Title:             title,
		BaseURL:           baseURL,
		RepositoryOwner:   repositoryOwner,
		RepositoryName:    repositoryName,
		HelmRepositoryURL: helmRepositoryURL,
		HelmDirectoryName: helmDirectoryName,
		Namespace:         namespace,
	}, nil
}

// Update is update method for a project.
func (p *Projects) Update(tx *sql.Tx, id int, baseURL, helmDirectoryName, namespace string) error {
	_, err := tx.Exec("UPDATE projects SET base_url = $1, helm_directory_name = $2, namespace = $3) WHERE id = $3", baseURL, helmDirectoryName, namespace, id)
	return err
}
