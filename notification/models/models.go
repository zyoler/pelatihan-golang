package models

import "time"

type Game struct {
	Id        int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama      string `json:"nama"`
	JenisItem string `json:"jenis_item"`
}

type Users struct {
	Id       int       `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama     string    `json:"nama"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type Toko struct {
	Id       int       `gorm:"primaryKey;autoIncrement;" json:"id"`
	UsersId  int       `json:"id_user"`
	NamaToko string    `json:"nama_toko"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type TokoDetail struct {
	Id         int `gorm:"primaryKey;autoIncrement;" json:"id"`
	TokoId     int `json:"id_toko"`
	GameId     int `json:"id_game"`
	JumlahItem int `json:"jumlah_Item"`
	HargaJual  int `json:"harga_jual"`
}

type TokoDetailItem struct {
	Toko       Toko
	TokoDetail TokoDetail
}

type Transaksi struct {
	Id      int `gorm:"primaryKey;autoIncrement;" json:"id"`
	GameId  int `json:"id_game"`
	UsersId int `json:"id_user"`
	TokoId  int `json:"id_toko"`
}

type DetailTransaksi struct {
	TransaksiId    int    `json:"id_transaksi"`
	JumlahItem     int    `json:"jumlah_item"`
	HargaTransaksi int    `json:"harga_transaksi"`
	Keterangan     string `json:"keterangan"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
