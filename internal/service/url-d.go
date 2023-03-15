package service

import (
	"awesomeProject-URL/internal/model"
	"awesomeProject-URL/internal/repository"
	"log"
)

type UrlDBService struct {
	repository repository.Url
}

func NewUrlDBService(repository repository.Url) *UrlDBService {
	return &UrlDBService{repository: repository}
}
func (s *UrlDBService) CreateUrl(url model.URL) (string, error) {

	// проверяем, есть ли url в БД
	fUrl, ok := s.repository.ContainsUrl(url)
	if ok {
		log.Printf("[S] Short url found in DB during creation: %s\n", fUrl.ShortUrl)
		return fullShortUrl(fUrl.ShortUrl), nil
	}
	// если нет - генерируем короткий url
	url.ShortUrl = genShortUrl()
	// записываем в БД
	err := s.repository.CreateUrl(url)
	return fullShortUrl(url.ShortUrl), err
}

func (s *UrlDBService) GetUrl(shortUrl string) (string, error) {
	url, err := s.repository.GetUrl(shortUrl)
	log.Printf("[S] Get url: %s\n", url.Url)
	return url.Url, err
}
