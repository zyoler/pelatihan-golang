package repository

import (
	"pemesanan/models"

	"gorm.io/gorm"
)

type Repo struct {
	app *gorm.DB
}

type PemesananRepoInterface interface {
	FindAllKategori() ([]models.Kategori, error)
	FindAllKereta() ([]models.Kereta, error)
	FindAllDetail() ([]models.DetailKereta, error)
	FindOneKereta(i interface{}, where map[string]interface{}) error
}

func NewRepo(app *gorm.DB) PemesananRepoInterface {
	return &Repo{app}
}
