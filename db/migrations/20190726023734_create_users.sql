
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS users (
id SERIAL PRIMARY KEY,
email varchar(255) NOT NULL,
oauth_token varchar(255) NOT NULL,
uuid bigint NOT NULL,
identifier varchar(255) NOT NULL,
avatar_url varchar(255) DEFAULT NULL,
created_at timestamp NOT NULL DEFAULT current_timestamp,
updated_at timestamp NOT NULL DEFAULT current_timestamp);

CREATE UNIQUE INDEX index_on_email ON users (email);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX index_on_email;

DROP TABLE users;
