package usecase

import (
	"pemesanan/models"
)

func (us *Usecase) GetAllKategori() ([]models.Kategori, error) {
	data, err := us.query.FindAllKategori()
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (us *Usecase) GetAllKereta() ([]models.Kereta, error) {
	data, err := us.query.FindAllKereta()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *Usecase) GetAllDetail() ([]models.DetailKereta, error) {
	data, err := uc.query.FindAllDetail()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *Usecase) GetDetailById(id int) (models.Kereta, error) {
	var data models.Kereta
	err := uc.query.FindOneKereta(&data, map[string]interface{}{"id": id})
	if err != nil {
		return data, err
	}
	return data, nil
}
