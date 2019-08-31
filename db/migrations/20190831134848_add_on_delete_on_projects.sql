
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE project_values DROP CONSTRAINT project_values_project_id_fkey;
ALTER TABLE project_values ADD FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE project_values DROP CONSTRAINT project_values_project_id_fkey;
ALTER TABLE project_values ADD FOREIGN KEY (project_id) REFERENCES projects(id);
