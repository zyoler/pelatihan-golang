package controllers

import (
	"net/http"
	"pemesanan/usecase"
)

type BaseController struct {
	uc usecase.UsecasePemesananInterface
}

type BaseControllerInterface interface {
	GetDataKategori(w http.ResponseWriter, r *http.Request)
	GetDataKereta(w http.ResponseWriter, r *http.Request)
	GetDataDetail(w http.ResponseWriter, r *http.Request)
	GetDataDetailById(w http.ResponseWriter, r *http.Request)
}

func NewController(c usecase.UsecasePemesananInterface) BaseControllerInterface {
	return &BaseController{uc: c}
}
