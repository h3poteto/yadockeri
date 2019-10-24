
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP INDEX index_on_email;
CREATE UNIQUE INDEX unique_uuid_on_users on users (uuid);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX unique_uuid_on_users;
CREATE UNIQUE INDEX index_on_email on users (email);

