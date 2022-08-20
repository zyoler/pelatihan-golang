package models

type DetailKereta struct {
	ID         uint `gorm:"primary_key"`
	KeretaId   int  `json:"kereta_id"`
	Kereta     Kereta
	KategoriId int `json:"kategori_id"`
	Kategori   Kategori
}

type Kereta struct {
	Id    int    `gorm:"primary_key;auto_increment" json:"id"`
	Nama  string `json:"nama"`
	Harga uint64 `json:"harga"`
}

type Kategori struct {
	Id   int    `gorm:"primary_key;auto_increment" json:"id"`
	Nama string `json:"nama"`
}

type Transaksi struct {
	Id         int `gorm:"primary_key;auto_increment" json:"id"`
	UserId     int `json:"user_id"`
	KeretaId   int `json:"kereta_id"`
	Kereta     Kereta
	KategoriId int `json:"kategori_id"`
	Kategori   Kategori
	Total      uint64 `json:"total"`
	Tanggal    string `json:"tanggal";gorm:DateTime;auto_now_add`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
