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

func (r *UrlDB) CreateUrl(url model.URL) error {

	urlQuery := fmt.Sprintf("INSERT INTO %s (url, short_url) VALUES ($1, $2)", urlTable)
	_, err := r.db.Exec(urlQuery, url.Url, url.ShortUrl)
	return err
}

func (r *UrlDB) GetUrl(shortUrl string) (model.URL, error) {
	var url model.URL
	urlQuery := fmt.Sprintf("SELECT * FROM %s WHERE short_url LIKE $1", urlTable)
	err := r.db.Get(&url, urlQuery, shortUrl)
	return url, err
}

func (r *UrlDB) ContainsUrl(url model.URL) (model.URL, bool) {

	urlQuery := fmt.Sprintf("SELECT * FROM %s WHERE url LIKE $1", urlTable)
	err := r.db.Get(&url, urlQuery, url.Url)
	if err != nil {
		return model.URL{}, false
	}
	return url, true
}
