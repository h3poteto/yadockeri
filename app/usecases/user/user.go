package user

import (
	"context"
	"time"

	"github.com/google/go-github/v27/github"
	"github.com/h3poteto/yadockeri/app/repositories/users"
	"github.com/h3poteto/yadockeri/config"
	"github.com/h3poteto/yadockeri/db"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type User struct {
	ID         int
	Email      string
	OauthToken string
	UUID       int64
	Identifier string
	AvatarURL  string
}

func FindOrCreateUser(token string) (int, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	githubUser, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return 0, errors.Wrap(err, "github api error")
	}

	// TODO: orgの指定がない場合はすべてを通す
	isMember, _, err := client.Organizations.IsMember(ctx, config.Element("organization"), githubUser.GetLogin())

	if err != nil {
		return 0, errors.Wrap(err, "github api error")
	}

	if !isMember {
		return 0, errors.New("not member")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	emails, _, _ := client.Users.ListEmails(ctx, nil)
	var primaryEmail string
	for _, email := range emails {
		if *email.Primary {
			primaryEmail = *email.Email
		}
	}

	u := users.New(db.SharedInstance().Connection)
	user, err := u.GetByUUID(*githubUser.ID)
	if err != nil {
		return u.Create(primaryEmail, token, *githubUser.ID, *githubUser.Login, *githubUser.AvatarURL)
	}
	if user.ID == 0 {
		return 0, errors.New("cannot find user by email")
	}
	if user.OauthToken != token {
		return u.Update(user.ID, primaryEmail, token, *githubUser.ID, *githubUser.Login, *githubUser.AvatarURL)
	}
	return user.ID, nil
}

func LoginConfirm(currentID int) (*User, error) {
	u := users.New(db.SharedInstance().Connection)
	user, err := u.GetByID(currentID)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("can not find user")
	}
	return &User{
		ID:         user.ID,
		Email:      user.Email,
		OauthToken: user.OauthToken,
		UUID:       user.UUID,
		Identifier: user.Identifier,
		AvatarURL:  user.AvatarURL,
	}, nil
}
