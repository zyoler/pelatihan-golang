package models

import "gorm.io/gorm"

type Notif struct {
	gorm.Model
	MhsId int    `json:"mhs_id"`
	Notif string `json:"notif"`
}
