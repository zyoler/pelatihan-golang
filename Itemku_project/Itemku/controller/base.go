package controller

import (
	"Itemku/usecase"
	"net/http"
)

type ctrl struct {
	us usecase.Usecase
}
type ControllerInterface interface {
	GetDataGame(w http.ResponseWriter, r *http.Request)

	GetDataUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	PostDataUser(w http.ResponseWriter, r *http.Request)
	UpdateDataUser(w http.ResponseWriter, r *http.Request)
	DeleteDataUser(w http.ResponseWriter, r *http.Request)

	GetDataToko(w http.ResponseWriter, r *http.Request)
	PostDataToko(w http.ResponseWriter, r *http.Request)
	UpdateDataToko(w http.ResponseWriter, r *http.Request)
	DeleteDataToko(w http.ResponseWriter, r *http.Request)

	PostDataTokoDetail(w http.ResponseWriter, r *http.Request)
	DeleteDataTokoDetail(w http.ResponseWriter, r *http.Request)

	GetDataTokoDetailById(w http.ResponseWriter, r *http.Request)

	Notif(w http.ResponseWriter, r *http.Request)
}

func NewController(us usecase.Usecase) ControllerInterface {
	return &ctrl{us: us}
}
