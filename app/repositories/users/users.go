package users

import (
	"database/sql"

	"github.com/h3poteto/yadockeri/app/domains/user"
)

// Users is repository struct for users table.
type Users struct {
	db *sql.DB
}

// New return a users repository.
func New(db *sql.DB) *Users {
	return &Users{
		db,
	}
}

// Create creates a users record, and returns the id.
func (u *Users) Create(email, oauthToken string, uuid int64, identifier, avatarURL string) (int, error) {
	var id int
	err := u.db.QueryRow("INSERT INTO users (email, oauth_token, uuid, identifier, avatar_url) VALUES ($1, $2, $3, $4, $5) RETURNING id", email, oauthToken, uuid, identifier, avatarURL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update updates a users record, and returns the id.
func (u *Users) Update(id int, email, oauthToken string, uuid int64, identifier, avatarURL string) (int, error) {
	_, err := u.db.Exec("UPDATE users SET email = $1, oauth_token = $2, uuid = $3, identifier = $4, avatar_url = $5 WHERE id = $6", email, oauthToken, uuid, identifier, avatarURL, id)
	return id, err
}

// GetByUUID find a user using uuid which is assign from GitHub.
func (u *Users) GetByUUID(uuid int64) (*user.User, error) {
	var id int
	var email, oauthToken, identifier, avatarURL string
	var findUUID int64
	err := u.db.QueryRow("SELECT id, email, oauth_token, uuid, identifier, avatar_url FROM users where uuid = $1", uuid).Scan(&id, &email, &oauthToken, &findUUID, &identifier, &avatarURL)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:         id,
		Email:      email,
		OauthToken: oauthToken,
		UUID:       findUUID,
		Identifier: identifier,
		AvatarURL:  avatarURL,
	}, nil
}

// GetByID find a user using id.
func (u *Users) GetByID(id int) (*user.User, error) {
	var findID int
	var email, oauthToken, identifier, avatarURL string
	var uuid int64
	err := u.db.QueryRow("SELECT id, email, oauth_token, uuid, identifier, avatar_url FROM users where id = $1", id).Scan(&findID, &email, &oauthToken, &uuid, &identifier, &avatarURL)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:         findID,
		Email:      email,
		OauthToken: oauthToken,
		UUID:       uuid,
		Identifier: identifier,
		AvatarURL:  avatarURL,
	}, nil
}
