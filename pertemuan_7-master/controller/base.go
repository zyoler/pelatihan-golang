package controller

import (
	"net/http"
	"pertemuan_6/usecase"
)

type ctrl struct {
	us usecase.Usecase
}
type ControllerInterface interface {
	GetData(w http.ResponseWriter, r *http.Request)
	PostData(w http.ResponseWriter, r *http.Request)
	UpdateData(w http.ResponseWriter, r *http.Request)
	DeleteData(w http.ResponseWriter, r *http.Request)
}

func NewController(us usecase.Usecase) ControllerInterface {
	return &ctrl{us: us}
}
