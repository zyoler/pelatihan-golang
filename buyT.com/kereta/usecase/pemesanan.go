package usecase

import "kereta/models"

func (r *Uc) PemesananKereta(stasiunA int, stasiunT int, ktg int) ([]models.Stasiun, []models.Kategori, error) {
	var stasiun []models.Stasiun
	var kategori []models.Kategori
	err := r.query.FindOneId(&stasiun, &kategori, stasiunA, stasiunT, ktg)
	if err != nil {
		return stasiun, kategori, err
	}
	return stasiun, kategori, nil
}

// func (r *Uc) FindAllPemesananByAll(Id string) (data models.Pemesanan, err error) {
// 	return r.query.FindAllPemesananByAll(Id)
// }

func (r *Uc) FindAllPemesananById(id string) (models.Pemesanan, error) {
	var data models.Pemesanan
	err := r.query.FindAllPemesananByAll(&data, id)
	if err != nil {
		return data, err
	}
	return data, nil
}
