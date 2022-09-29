package getfile

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
	payload := NewUGetFileRequest(req)
	file, err := i.uploadRepo.GetFile(ctx, payload)
	if err != nil {
		return InportResponse{}, err
	}

	return InportResponse{
		File: file,
	}, nil
}
