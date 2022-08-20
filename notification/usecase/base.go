package usecase

import (
	"notification/models"
	"notification/repository"
)

type UC struct {
	queryrepo repository.Repo
}
type Usecase interface {
	GetDataUser() ([]models.Users, error)
	InsertDataUser(models.Users) error
	UpdateDataUser(models.Users) error
	DeleteDataUser(id int) error
	NotifUser(id int) (models.NotifItemku, error)
}

func NewUsecase(r repository.Repo) Usecase {
	return &UC{queryrepo: r}
}
