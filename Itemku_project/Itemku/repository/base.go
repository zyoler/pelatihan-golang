package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type repo struct {
	apps *gorm.DB
}
type Repo interface {
	FindAll(i interface{}) error
	InsertData(i interface{}) error
	UpdateData(i interface{}, where map[string]interface{}, data map[string]interface{}) error
	DeleteData(i interface{}, where map[string]interface{}) error
	DinamicFindQueryRaw(i interface{}, query string) (*sql.Rows, error)
}

func NewRepo(apps *gorm.DB) Repo {
	return &repo{apps}
}
