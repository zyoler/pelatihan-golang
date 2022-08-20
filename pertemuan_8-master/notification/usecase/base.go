package usecase

import (
	"notification/models"
	"notification/repository"
)

type UC struct {
	queryrepo repository.Repo
}
type Usecase interface {
	GetDataMHS() ([]models.MhsPub, error)
	InsertDataMHS(models.MhsPub) error
	UpdateDataMhs(models.MhsPub) error
	DeleteDataMhs(id int) error
	NotifMhs(id int) (models.Notif, error)
}

func NewUsecase(r repository.Repo) Usecase {
	return &UC{queryrepo: r}
}
