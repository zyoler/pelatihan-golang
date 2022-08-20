package models

type MhsPub struct {
	Id   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
