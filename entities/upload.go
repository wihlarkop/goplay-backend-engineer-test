package entities

type UploadFile struct {
	Id        int    `json:"id" db:"id"`
	ImagePath string `json:"image_path" db:"path"`
	UploadBy  string `json:"upload_by" db:"upload_by"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type UploadFileFilter struct {
	Page  int
	Limit int
}
