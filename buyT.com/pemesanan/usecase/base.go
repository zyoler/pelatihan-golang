package usecase

import (
	"pemesanan/models"
	"pemesanan/repository"
)

type Usecase struct {
	query repository.PemesananRepoInterface
}

type UsecasePemesananInterface interface {
	GetAllKategori() ([]models.Kategori, error)
	GetAllKereta() ([]models.Kereta, error)
	GetAllDetail() ([]models.DetailKereta, error)
	GetDetailById(id int) (models.Kereta, error)
}

func NewUsecase(r repository.PemesananRepoInterface) UsecasePemesananInterface {
	return &Usecase{query: r}
}
