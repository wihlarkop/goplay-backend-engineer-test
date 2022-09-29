package file

import (
	"context"

	"goplay-backend-engineer-test/entities"
	"goplay-backend-engineer-test/helper"
)

//go:generate mockgen -destination=mock/uom.go -package=mock repositories/upload IRepo
type IRepo interface {
	GetFiles(ctx context.Context, req entities.UploadFileFilter) ([]entities.UploadFile, helper.MetaTpl, error)
	GetFile(ctx context.Context, req entities.UploadFile) (entities.UploadFile, error)
	CreateFile(ctx context.Context, req entities.UploadFile) (entities.UploadFile, error)
}
