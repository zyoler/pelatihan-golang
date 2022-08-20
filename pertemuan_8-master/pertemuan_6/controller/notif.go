package controller

import (
	"log"
	"net/http"
	"pertemuan_6/models"
	"pertemuan_6/tools"
)

func (c *ctrl) Notif(w http.ResponseWriter, r *http.Request) {

	method := "GET"
	baseUrl := "http://localhost:5001/get-notification/1"
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
