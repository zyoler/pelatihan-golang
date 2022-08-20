package controller

import (
	"Itemku/models"
	"Itemku/tools"
	"log"
	"net/http"
)

func (c *ctrl) Notif(w http.ResponseWriter, r *http.Request) {

	method := "GET"
	baseUrl := "http://localhost:8001/get-notification/1"
	header := map[string][]string{
		"Content-type": {"application/json"},
	}
	var ResponseServiceNotif models.ResponseNotifService

	err := tools.Curl(method, baseUrl, nil, &ResponseServiceNotif, header)
	if err != nil {
		log.Println("Error Hit Service Notification=> ", err)
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ResponseApi(w, 200, ResponseServiceNotif.Notif, "Get Data Success")
}
