package models

import "gorm.io/gorm"

type ResponseService struct {
	code    int    `json:"code"`
	messsge string `json:"message"`
}

type Notif struct {
	gorm.Model
	MhsId int    `json:"mhs_id"`
	Notif string `json:"notif"`
}

type ResponseNotifService struct {
	ResponseService
	Notif `json:"data"`
}

type RequestNotify struct {
	IdPemesanan int         `json:"pemesanan_id"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
