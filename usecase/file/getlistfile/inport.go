package getlistfile

import (
	"context"

	"goplay-backend-engineer-test/entities"
	"goplay-backend-engineer-test/helper"
)

//go:generate mockgen -destination=mock/inport.go -package=mock usecases/upload/getlistfile Inport
type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	entities.UploadFileFilter
}

type InportResponse struct {
	File       []entities.UploadFile `json:"file"`
	Pagination helper.MetaTpl        `json:"pagination"`
}
