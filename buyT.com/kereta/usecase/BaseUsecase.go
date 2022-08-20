package usecase

import (
	"kereta/models"
	"kereta/repository"
)

type Uc struct {
	query repository.RepositoryInterface
}

type UsecaseInterface interface {
	Stasiun() ([]models.Stasiun, error)
	GetFirstStasiun(id string) (models.Stasiun, error)
	GetAllStasiunByKota(id string) ([]models.Stasiun, error)
	Register(data models.User) (models.User, error)
	// PemesananKereta(stasiunA int, stasiunT int, ktg int) ([]models.Stasiun, []models.Kategori, error)
	FindAllPemesananById(Id string) (models.Pemesanan, error)
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &Uc{query: r}
}
