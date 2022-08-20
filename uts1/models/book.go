package models

type Books struct {
	Id          int    `gorm:"primaryKey;autoIncrement;"`
	Judul       string `json:"judul"`
	Penerbit    string `json:"penerbit"`
	TahunTerbit int    `json:"tahunterbit"`
	Harga       int    `json:"harga"`
}
