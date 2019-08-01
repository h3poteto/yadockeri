package project_values

import (
	"database/sql"

	"github.com/h3poteto/yadockeri/app/values"
)

type ProjectValues struct {
	db *sql.DB
}

func New(db *sql.DB) *ProjectValues {
	return &ProjectValues{
		db,
	}
}

func (p *ProjectValues) Create(tx *sql.Tx, projectID int, key, value string) (int, error) {
	var id int
	err := tx.QueryRow("INSERT INTO project_values (project_id, key, override_value) VALUES ($1, $2, $3) RETURNING id", projectID, key, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *ProjectValues) GetByProject(projectID int) ([]*values.OverrideValue, error) {
	var overrides []*values.OverrideValue
	rows, err := p.db.Query("SELECT id, project_id, key, override_value FROM project_values WHERE project_id = $1", projectID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, findProjectID int
		var key, value string
		err = rows.Scan(&id, &findProjectID, &key, &value)
		if err != nil {
			return nil, err
		}
		v := &values.OverrideValue{
			Key:   key,
			Value: value,
		}
		overrides = append(overrides, v)
	}
	return overrides, nil
}
