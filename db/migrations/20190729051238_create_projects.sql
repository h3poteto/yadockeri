
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS projects (
id SERIAL PRIMARY KEY,
user_id int REFERENCES users(id),
title varchar(255) NOT NULL,
base_url varchar(255) NOT NULL,
repository_owner varchar(255) NOT NULL,
repository_name varchar(255) NOT NULL,
helm_repository_url varchar(255) NOT NULL,
helm_directory_name varchar(255) NOT NULL,
namespace varchar(255) NOT NULL,
created_at timestamp NOT NULL DEFAULT current_timestamp,
updated_at timestamp NOT NULL DEFAULT current_timestamp);

CREATE UNIQUE INDEX repository_owner_and_repository_name_on_projects ON projects (repository_owner, repository_name);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX repository_owner_and_repository_name_on_projects;

DROP TABLE projects;
