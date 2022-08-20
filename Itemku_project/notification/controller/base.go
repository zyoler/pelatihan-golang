package controller

import (
	"notification/usecase"
)

type ctrl struct {
	us usecase.Usecase
}
type ControllerInterface interface {
}

func NewController(us usecase.Usecase) ControllerInterface {
	return &ctrl{us: us}
}
