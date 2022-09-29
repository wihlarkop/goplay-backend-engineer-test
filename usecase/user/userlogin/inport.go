package userlogin

import (
	"context"

	"goplay-backend-engineer-test/entities"
)

//go:generate mockgen -destination=mock/inport.go -package=mock usecases/user/login Inport
type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InportResponse struct {
	entities.User
	JwtToken string `json:"token"`
}

func NewLoginRequest(req InportRequest) entities.User {
	return entities.User{
		Username: req.Username,
		Password: req.Password,
	}
}
