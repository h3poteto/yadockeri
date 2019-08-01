
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS branches (
id SERIAL PRIMARY KEY,
project_id int REFERENCES projects(id),
user_id int REFERENCES users(id),
name varchar(255) NOT NULL,
url varchar(255) DEFAULT NULL,
created_at timestamp NOT NULL DEFAULT current_timestamp,
updated_at timestamp NOT NULL DEFAULT current_timestamp);

CREATE UNIQUE INDEX project_id_and_name_on_branches ON branches (project_id, name);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX project_id_and_name_on_branches;

DROP TABLE branches;
