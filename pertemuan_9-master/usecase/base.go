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
	InsertDataMHS(models.MhsPub) error
	UpdateDataMhs(models.MhsPub) error
	DeleteDataMhs(id int) error
	GetDataDetailMhs(id string) (error, *[]models.NilaiMhsData)
}

func NewUsecase(r repository.Repo) Usecase {
	return &UC{queryrepo: r}
}
