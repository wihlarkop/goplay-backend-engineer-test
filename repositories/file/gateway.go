package file

import "github.com/jmoiron/sqlx"

type repo struct {
	sqlite *sqlx.DB
}

func NewRepo(sqlite *sqlx.DB) IRepo {
	return &repo{
		sqlite: sqlite,
	}
}
