package usecase

import (
	"Itemku/models"
	"time"
)

func (r *UC) GetDataToko() ([]models.Toko, error) {
	var Model []models.Toko
	err := r.queryrepo.FindAll(&Model)
	if err != nil {
		return Model, err
	}
	return Model, nil
}

func (r *UC) InsertDataToko(data models.Toko) error {
	err := r.queryrepo.InsertData(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) DeleteDataToko(id int) error {
	Where := make(map[string]interface{})
	Where["id"] = id
	var Table models.Toko
	err := r.queryrepo.DeleteData(&Table, Where)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) UpdateDataToko(data models.Toko) error {
	Where := make(map[string]interface{})
	Where["id"] = data.Id

	dataUpdates := make(map[string]interface{})
	dataUpdates["nama"] = data.NamaToko
	dataUpdates["update_at"] = time.Now()

	err := r.queryrepo.UpdateData(&data, Where, dataUpdates)
	if err != nil {
		return err
	}
	return nil
}
