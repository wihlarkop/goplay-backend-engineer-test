package userlogin

import (
	"context"

	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/repositories/user"
)

type interactor struct {
	userRepo user.IRepo
}

func NewUsecase(userRepo user.IRepo) Inport {
	return interactor{
		userRepo: userRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) (InportResponse, error) {
	payload := NewLoginRequest(req)
	user, err := i.userRepo.GetUser(ctx, payload)
	if err != nil {
		return InportResponse{}, err
	}

	request := helper.TokenRequest{
		Name:      user.Username,
		CreatedAt: user.CreatedAt,
	}

	token, err := helper.GenerateToken(request)
	if err != nil {
		return InportResponse{}, err
	}

	return InportResponse{
		User:     user,
		JwtToken: token,
	}, nil
}
