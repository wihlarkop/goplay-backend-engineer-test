package getlistfile

import (
	"context"

	"goplay-backend-engineer-test/repositories/file"
)

type interactor struct {
	uploadRepo file.IRepo
}

func NewUsecase(uploadRepo file.IRepo) Inport {
	return interactor{
		uploadRepo: uploadRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) (InportResponse, error) {
	files, pagination, err := i.uploadRepo.GetFiles(ctx, req.UploadFileFilter)
	if err != nil {
		return InportResponse{}, err
	}

	return InportResponse{
		File:       files,
		Pagination: pagination,
	}, nil
}
