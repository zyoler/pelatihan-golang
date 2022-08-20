package repository

import (
	"kereta/models"

	"gorm.io/gorm"
)

type Repo struct {
	app *gorm.DB
}

type RepositoryInterface interface {
	FindAll() ([]models.Stasiun, error)
	FindOneId(i interface{}, args ...interface{}) error
	FindAllPemesananByAll(i interface{}, args ...interface{}) error
	FindAllKotaId(i string) ([]models.Stasiun, error)
	// PemesananAdd(stasiunA int, stasiunT int, ktg int) ([]models.Stasiun, []models.Kategori, error)
	Create(i interface{}) error
	// InsertData(i interface{}) error
	// UpdateData(i interface{}, where map[string]interface{}, data map[string]interface{}) error
	// DeleteData(i interface{}, where map[string]interface{}) error
	// DinamicData(i interface{}, query string) (*sql.Rows, error)
}

func NewRepository(app *gorm.DB) RepositoryInterface {
	return &Repo{app}
}
