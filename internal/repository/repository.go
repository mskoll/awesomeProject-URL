package repository

import (
	"awesomeProject-URL/internal/model"
	"github.com/jmoiron/sqlx"
)

type Url interface {
	CreateUrl(url model.URL) error
	GetUrl(shortUrl string) (model.URL, error)
	ContainsUrl(url model.URL) (model.URL, bool)
}

type Repository struct {
	Url
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Url: NewUrlDB(db)}
}
