package models

type MhsPub struct {
	Id   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
	// NilaiMhsPub NilaiMhsPub `gorm:"-"`
}

type NilaiMhsPub struct {
	IdNilai int    `gorm:"primaryKey;autoIncrement;" json:"id_nilai"`
	Id      int    `json:"id_mhs"`
	Matkul  string `json:"matkul"`
	Nilai   int    `json:"nilai"`
}

type NilaiMhsData struct {
	MhsPub
	NilaiMhsPub
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
