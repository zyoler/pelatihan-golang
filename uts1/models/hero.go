package models

type HeroML struct {
	Id   int    `gorm:"primaryKey;autoIncrement;"`
	Nama string `json:"nama"`
	Role string `json:"role"`
}
