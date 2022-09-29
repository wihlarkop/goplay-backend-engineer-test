package getfile

import (
	"context"

	"goplay-backend-engineer-test/entities"
)

//go:generate mockgen -destination=mock/inport.go -package=mock usecases/upload/getfile Inport
type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	entities.UploadFile
}

type InportResponse struct {
	File entities.UploadFile `json:"file"`
}

func NewUGetFileRequest(req InportRequest) entities.UploadFile {
	return entities.UploadFile{
		Id: req.Id,
	}
}
