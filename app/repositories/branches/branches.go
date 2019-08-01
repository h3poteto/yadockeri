package branches

import (
	"database/sql"

	"github.com/h3poteto/yadockeri/app/domains/branch"
)

type Branches struct {
	db *sql.DB
}

func New(db *sql.DB) *Branches {
	return &Branches{
		db,
	}
}

func (b *Branches) GetByProject(projectID int) ([]*branch.Branch, error) {
	var branches []*branch.Branch
	rows, err := b.db.Query("SELECT id, project_id, user_id, name, url FROM branches WHERE project_id = $1", projectID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, findProjectID, userID int
		var name string
		var url sql.NullString
		err = rows.Scan(&id, &findProjectID, &userID, &name, &url)
		if err != nil {
			return nil, err
		}
		b := &branch.Branch{
			ID:        id,
			ProjectID: findProjectID,
			UserID:    userID,
			Name:      name,
			URL:       url.String,
		}
		branches = append(branches, b)
	}
	return branches, nil
}

func (b *Branches) Create(projectID, userID int, name string) (int, error) {
	var id int
	err := b.db.QueryRow("INSERT INTO branches (project_id, user_id, name) VALUES ($1, $2, $3) RETURNING id", projectID, userID, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *Branches) GetByID(id int) (*branch.Branch, error) {
	var findID, projectID, userID int
	var name string
	var url sql.NullString
	err := b.db.QueryRow("SELECT id, project_id, user_id, name, url FROM branches WHERE id = $1", id).Scan(&findID, &projectID, &userID, &name, &url)
	if err != nil {
		return nil, err
	}
	return &branch.Branch{
		ID:        findID,
		ProjectID: projectID,
		UserID:    userID,
		Name:      name,
		URL:       url.String,
	}, nil
}

func (b *Branches) DeleteByID(id int) error {
	_, err := b.db.Exec("DELETE FROM branches WHERE id = $1", id)
	return err
}

func (b *Branches) Update(branch *branch.Branch) error {
	_, err := b.db.Exec("UPDATE branches SET project_id = $2, user_id = $3, name = $4, url = $5 WHERE id = $1", branch.ID, branch.ProjectID, branch.UserID, branch.Name, branch.URL)
	return err
}
