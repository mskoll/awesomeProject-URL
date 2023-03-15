package service

import (
	"awesomeProject-URL/internal/repository"
	"github.com/spf13/viper"
	"log"
)

type Provider struct {
	Url
}

func NewUrlService(repository *repository.Repository) *Provider {
	// проверяем флаг запуска, отдаем нужный сервис в зависимости от флага
	if flag := viper.GetBool("flagDB"); flag {
		log.Printf("[S] Service DB\n")
		return &Provider{Url: NewUrlDBService(repository.Url)}
	} else {
		log.Printf("[S] Service cache\n")
		return &Provider{Url: NewUrlCacheService(repository.Url)}
	}
}
