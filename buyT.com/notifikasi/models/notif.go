package models

import (
	"gorm.io/gorm"
)

type Notif struct {
	Id          int    `json:"id"`
	IdPemesanan int    `json:"pemesanan_id"`
	Request     string `json:"request"`
	Sent        bool   `json:"sent"`
	gorm.Model
}

type RequestNotify struct {
	IdPemesanan int         `json:"pemesanan_id"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
