package service

import (
	"awesomeProject-URL/internal/model"
	"awesomeProject-URL/internal/repository"
	"log"
)

// UrlCacheService структура сервиса с кэшем
// cacheShort для поиска коротких url для get
// cacheFull для проверки, есть ли уже url для create
type UrlCacheService struct {
	repository repository.Url
	cacheShort map[string]string
	cacheFull  map[string]string
}

func NewUrlCacheService(repository repository.Url) *UrlCacheService {
	return &UrlCacheService{repository: repository,
		cacheShort: make(map[string]string),
		cacheFull:  make(map[string]string)}
}

func (s *UrlCacheService) CreateUrl(url model.URL) (string, error) {
	// проверяем, есть ли url в кэше
	shortUrl, ok := s.cacheFull[url.Url]
	if ok {
		log.Printf("[S] Short url found in cache during creation: %s\n", shortUrl)
		return fullShortUrl(shortUrl), nil
	}
	// проверяем, есть ли url в БД
	fUrl, ok := s.repository.ContainsUrl(url)
	if ok {
		log.Printf("[S] Short url found in DB during creation: %s\n", fUrl.ShortUrl)
		// записываем в кэш на будущее
		s.cacheShort[fUrl.ShortUrl] = fUrl.Url
		s.cacheFull[fUrl.Url] = fUrl.ShortUrl
		return fullShortUrl(fUrl.ShortUrl), nil
	}
	// если нет - генерируем короткий url
	url.ShortUrl = genShortUrl()
	// записываем в БД
	err := s.repository.CreateUrl(url)
	// записываем в кэш
	s.cacheShort[url.ShortUrl] = url.Url
	s.cacheFull[url.Url] = url.ShortUrl
	return fullShortUrl(url.ShortUrl), err
}

func (s *UrlCacheService) GetUrl(shortUrl string) (string, error) {

	// смотрим url в кэше
	url, ok := s.cacheShort[shortUrl]
	if !ok {
		// если в кэше нет - смотрим в БД
		url, err := s.repository.GetUrl(shortUrl)
		if err != nil {
			return "", err
		}
		// пишем в кэш на будущее
		s.cacheShort[shortUrl] = url.Url
		s.cacheFull[url.Url] = shortUrl
		log.Printf("[S] Get url from DB: %s\n", url.Url)
		return url.Url, nil
	}
	log.Printf("[S] Get url from cache: %s\n", url)
	return url, nil
}
