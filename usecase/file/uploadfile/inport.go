package uploadfile

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"goplay-backend-engineer-test/entities"
)

//go:generate mockgen -destination=mock/inport.go -package=mock goplay-backend-engineer-test/usecase/file/uploadfile Inport
type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	Location string               `json:"location" form:"location" validate:"required"`
	File     multipart.FileHeader `json:"file" form:"file" validate:"required"`
	UploadBy string               `json:"uploaded_by"`
}

type InportResponse struct {
	Url       string `json:"url"`
	UploadBy  string `json:"upload_by"`
	CreatedAt string `json:"created_at"`
}

func NewUploadFileRequest(req InportRequest) entities.UploadFile {
	return entities.UploadFile{
		ImagePath: req.Location,
		UploadBy:  req.UploadBy,
		CreatedAt: time.Now().String(),
	}
}

func NewUploadFileResponse(req entities.UploadFile) InportResponse {
	filepath := fmt.Sprintf("http://localhost:3000/%s", req.ImagePath)
	return InportResponse{
		Url:       filepath,
		UploadBy:  req.UploadBy,
		CreatedAt: req.CreatedAt,
	}
}
