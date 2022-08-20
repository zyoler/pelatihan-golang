package repository

import "pemesanan/models"

func (r *Repo) FindAllKategori() ([]models.Kategori, error) {
	var kategoris []models.Kategori
	err := r.app.Find(&kategoris).Error
	if err != nil {
		return nil, err
	}
	return kategoris, err
}

func (r *Repo) FindAllKereta() ([]models.Kereta, error) {
	var Kereta []models.Kereta
	err := r.app.Find(&Kereta).Error
	if err != nil {
		return nil, err
	}
	return Kereta, err
}

func (r *Repo) FindAllDetail() ([]models.DetailKereta, error) {
	var DetailKereta []models.DetailKereta
	err := r.app.Joins("Kereta").Joins("Kategori").Find(&DetailKereta).Error
	if err != nil {
		return nil, err
	}
	return DetailKereta, err
}

func (r *Repo) FindOneKereta(i interface{}, where map[string]interface{}) error {
	result := r.app.Find(i, where)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
