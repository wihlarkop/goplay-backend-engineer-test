package user

import (
	"context"

	"goplay-backend-engineer-test/entities"
)

//go:generate mockgen -destination=mock/user.go -package=mock repositories/user IRepo
type IRepo interface {
	GetUser(ctx context.Context, req entities.User) (entities.User, error)
}
