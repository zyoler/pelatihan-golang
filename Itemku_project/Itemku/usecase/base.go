package usecase

import (
	"Itemku/models"
	"Itemku/repository"
)

type UC struct {
	queryrepo repository.Repo
}

type Usecase interface {
	GetDataGame() ([]models.Game, error)

	GetDataUser() ([]models.Users, error)
	GetUserById(id string) (*[]models.Users, error)
	InsertDataUser(models.Users) error
	UpdateDataUser(models.Users) error
	DeleteDataUser(id int) error

	GetDataToko() ([]models.Toko, error)
	InsertDataToko(models.Toko) error
	UpdateDataToko(models.Toko) error
	DeleteDataToko(id int) error

	InsertDataTokoDetail(models.TokoDetail) error
	DeleteDataTokoDetail(id int) error

	GetDataTokoDetailById(id string) (*[]models.TokoDetailItem, error)
}

func NewUsecase(r repository.Repo) Usecase {
	return &UC{queryrepo: r}
}
