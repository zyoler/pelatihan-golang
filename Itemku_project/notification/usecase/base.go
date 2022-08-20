package usecase

import (
	"notification/models"
	"notification/repository"
)

type UC struct {
	queryrepo repository.Repo
}

type Usecase interface {
	NotifUser(id int) (models.NotifItemku, error)
}

func NewUsecase(r repository.Repo) Usecase {
	return &UC{queryrepo: r}
}
