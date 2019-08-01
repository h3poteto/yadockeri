package users

import (
	"database/sql"

	"github.com/h3poteto/yadockeri/app/domains/user"
)

type Users struct {
	db *sql.DB
}

func New(db *sql.DB) *Users {
	return &Users{
		db,
	}
}

func (u *Users) Create(email, oauthToken string, uuid int64, identifier, avatarURL string) (int, error) {
	var id int
	err := u.db.QueryRow("INSERT INTO users (email, oauth_token, uuid, identifier, avatar_url) VALUES ($1, $2, $3, $4, $5) RETURNING id", email, oauthToken, uuid, identifier, avatarURL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *Users) GetByEmail(email string) (*user.User, error) {
	var id int
	var findEmail, oauthToken, identifier, avatarURL string
	var uuid int64
	err := u.db.QueryRow("SELECT id, email, oauth_token, uuid, identifier, avatar_url FROM users where email = $1", email).Scan(&id, &findEmail, &oauthToken, &uuid, &identifier, &avatarURL)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:         id,
		Email:      findEmail,
		OauthToken: oauthToken,
		UUID:       uuid,
		Identifier: identifier,
		AvatarURL:  avatarURL,
	}, nil
}

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
