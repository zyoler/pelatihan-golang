package usecase

import (
	"pertemuan_6/models"
	"pertemuan_6/repository"
)

type UC struct {
	queryrepo repository.Repo
}
type Usecase interface {
	GetDataMHS() ([]models.MhsPub, error)
}

func NewUsecase(r repository.Repo) Usecase {
	return &UC{queryrepo: r}
}
