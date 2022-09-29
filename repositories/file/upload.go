package file

import (
	"context"

	"goplay-backend-engineer-test/entities"
	"goplay-backend-engineer-test/helper"
)

func (r *repo) GetFiles(ctx context.Context, req entities.UploadFileFilter) ([]entities.UploadFile, helper.MetaTpl, error) {
	var (
		pagination helper.MetaTpl
		result     struct {
			Rows     []entities.UploadFile
			TotalRow int `json:"total_row"`
		}
	)

	page, limit, offset := helper.BuildPagination(req.Page, req.Limit)

	query := `select
					id,
					path,
					upload_by,
					created_at
			  from
					upload_file
			  limit $1 offset $2
	`

	err := r.sqlite.SelectContext(ctx, &result.Rows, query, limit, offset)
	if err != nil {
		return result.Rows, pagination, helper.ErrFatalQuery
	}

	pagination.Page = page
	pagination.Limit = limit
	pagination.TotalData = result.TotalRow

	return result.Rows, pagination, nil
}

func (r *repo) GetFile(ctx context.Context, req entities.UploadFile) (entities.UploadFile, error) {
	var files entities.UploadFile

	query := `select id, path, upload_by, created_at from upload_file where id = $1`

	err := r.sqlite.QueryRowxContext(ctx, query, req.Id).StructScan(&files)
	if err != nil {
		return files, helper.ErrFatalQuery
	}

	return files, nil
}

func (r *repo) CreateFile(ctx context.Context, req entities.UploadFile) (entities.UploadFile, error) {
	var result entities.UploadFile

	query := `
		INSERT INTO upload_file ( 
			path, upload_by, created_at
		) VALUES ($1, $2, $3) 
		RETURNING *
	`

	err := r.sqlite.QueryRowxContext(ctx, query, req.ImagePath, req.UploadBy, req.CreatedAt).StructScan(&result)
	if err != nil {
		return result, helper.ErrFatalQuery
	}

	return result, nil
}
