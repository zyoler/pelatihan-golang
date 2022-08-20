package models

import "gorm.io/gorm"

type Kategori struct {
	Id   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama string `json:"nama"`
	gorm.Model
}

type Product struct {
	Id    int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama  string `json:"nama"`
	Harga string `json:"harga"`
	gorm.Model
}

type User struct {
	Id   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
	gorm.Model
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
