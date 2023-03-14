package model

type URL struct {
	Id       int    `json:"-" db:"id"`
	Url      string `json:"url"`
	ShortUrl string `json:"short_url" db:"short_url"`
}
