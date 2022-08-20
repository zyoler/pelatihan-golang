package usecase

import "pertemuan_6/models"

func (r *UC) GetDataMHS() ([]models.MhsPub, error) {
	var Modelmhs []models.MhsPub
	err := r.queryrepo.FindAll(Modelmhs)
	if err != nil {
		return Modelmhs, err
	}
	return Modelmhs, nil
}
