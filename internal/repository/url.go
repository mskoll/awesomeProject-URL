package repository

import (
	"awesomeProject-URL/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UrlDB struct {
	db *sqlx.DB
}

func NewUrlDB(db *sqlx.DB) *UrlDB {
	return &UrlDB{db: db}
}

func (r *UrlDB) CreateUrl(url model.URL) (int, error) {
	var urlId int
	urlQuery := fmt.Sprintf("INSERT INTO %s (url, short_url) VALUES ($1, $2) RETURNING id", urlTable)

	row := r.db.QueryRow(urlQuery, url.Url, url.ShortUrl)
	if err := row.Scan(&urlId); err != nil {
		return 0, err
	}
	return urlId, nil
}

func (r *UrlDB) GetUrl(shortUrl string) (model.URL, error) {
	var url model.URL
	urlQuery := fmt.Sprintf("SELECT * FROM %s WHERE short_url = $1", urlTable)
	err := r.db.Get(&url, urlQuery, shortUrl)
	if err != nil {
		return url, err
	}
	return url, err
}
