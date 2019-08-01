
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS project_values(
id SERIAL PRIMARY KEY,
project_id int REFERENCES projects(id),
key varchar(255) NOT NULL,
override_value varchar(255) NOT NULL,
created_at timestamp NOT NULL DEFAULT current_timestamp,
update_dat timestamp NOT NULL DEFAULT current_timestamp);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE project_values;
