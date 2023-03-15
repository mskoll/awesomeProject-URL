package service

import (
	"awesomeProject-URL/internal/model"
	"awesomeProject-URL/internal/repository"
)

type Url interface {
	CreateUrl(url model.URL) (string, error)
	GetUrl(shortUrl string) (string, error)
}

type Service struct {
	Url
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Url: NewUrlService(repository)}
}
