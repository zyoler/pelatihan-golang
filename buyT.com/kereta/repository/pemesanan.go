package repository

import (
	"kereta/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.Connect()
}

func (r *Repo) FindAllPemesananByAll(i interface{}, args ...interface{}) error {
	result := r.app.Joins("Kereta").Joins("Kategori").First(i, args...)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
