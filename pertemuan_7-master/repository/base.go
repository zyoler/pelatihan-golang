package repository

import (
	"gorm.io/gorm"
)

type repo struct {
	apps *gorm.DB
}
type Repo interface {
	FindAll(i interface{}) error
	// InsertData(i interface{}) error
	// UpdateData(i interface{}, where map[string]interface{}) error
}

func NewRepo(apps *gorm.DB) Repo {
	return &repo{apps}
}
