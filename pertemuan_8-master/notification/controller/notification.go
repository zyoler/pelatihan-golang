package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (c *ctrl) NotifDataUser(w http.ResponseWriter, r *http.Request) {
	idmhs := chi.URLParam(r, "id")
	if idmhs == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idinteger, err := strconv.Atoi(idmhs)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	data, err := c.us.NotifMhs(idinteger)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	log.Println("Sukses Get Data Notif")
	ResponseApi(w, 200, data, "Get Notification Success")
	return
}
