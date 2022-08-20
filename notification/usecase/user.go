package usecase

import (
	"notification/models"
)

func (r *UC) GetDataUser() ([]models.Users, error) {
	var Modeluser []models.Users
	err := r.queryrepo.FindAll(Modeluser)
	if err != nil {
		return Modeluser, err
	}
	return Modeluser, nil
}

func (r *UC) InsertDataUser(data models.Users) error {
	err := r.queryrepo.InsertData(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) UpdateDataUser(data models.Users) error {
	var Where map[string]interface{}
	Where["id_user"] = data.IdUser

	var dataUpdates map[string]interface{}
	dataUpdates["nama"] = data.Nama
	dataUpdates["email"] = data.Email

	err := r.queryrepo.UpdateData(&data, Where, dataUpdates)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) DeleteDataUser(id int) error {
	var Where map[string]interface{}
	Where["id_user"] = id
	var TableUser models.Users
	err := r.queryrepo.DeleteData(&TableUser, Where)
	if err != nil {
		return err
	}
	return nil
}
