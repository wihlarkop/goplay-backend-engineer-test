package user

import (
	"context"

	"goplay-backend-engineer-test/entities"
	"goplay-backend-engineer-test/helper"
)

func (r *repo) GetUser(ctx context.Context, req entities.User) (entities.User, error) {
	var users []entities.User
	var user entities.User

	query := `select id, username, password, created_at from users where username = $1`

	err := r.sqlite.SelectContext(ctx, &users, query, req.Username)
	if err != nil {
		return user, helper.ErrFatalQuery
	}

	if len(users) > 0 {
		user = users[0]
	}

	return user, nil
}
