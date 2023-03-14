package service

import (
	"awesomeProject-URL/internal/model"
	"awesomeProject-URL/internal/repository"
)

type UrlService struct {
	repository repository.Url
}

func NewUrlService(repository repository.Url) *UrlService {
	return &UrlService{repository: repository}
}
func (s *UrlService) CreateUrl(url model.URL) (int, error) {

	urlId, err := s.repository.CreateUrl(url)

	return urlId, err
}

func (s *UrlService) GetUrl(shortUrl string) (model.URL, error) {
	return s.GetUrl(shortUrl)
}
