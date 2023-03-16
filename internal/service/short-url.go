package service

import (
	"github.com/spf13/viper"
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genShortUrl() string {
	n := 4
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func fullShortUrl(shortUrl string) string {
	sb := strings.Builder{}
	sb.Grow(35)
	sb.WriteString("http://localhost:")
	sb.WriteString(viper.GetString("port"))
	sb.WriteString("/")
	sb.WriteString(shortUrl)
	return sb.String()
}
