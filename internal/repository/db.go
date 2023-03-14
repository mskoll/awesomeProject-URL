package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	urlTable = "url"
)

type Conf struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// InitDB подключение к БД
func InitDB(conf Conf) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connToString(conf))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connToString(info Conf) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
