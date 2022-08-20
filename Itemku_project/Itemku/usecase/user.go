package usecase

import (
	"Itemku/models"
	"time"
)

func (r *UC) GetDataGame() ([]models.Game, error) {
	var Model []models.Game
	err := r.queryrepo.FindAll(&Model)
	if err != nil {
		return Model, err
	}
	return Model, nil
}

func (r *UC) GetDataUser() ([]models.Users, error) {
	var Modeluser []models.Users
	err := r.queryrepo.FindAll(&Modeluser)
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

func (r *UC) DeleteDataUser(id int) error {
	Where := make(map[string]interface{})
	Where["id"] = id
	var TableMhs models.Users
	err := r.queryrepo.DeleteData(&TableMhs, Where)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) UpdateDataUser(data models.Users) error {
	Where := make(map[string]interface{})
	Where["id"] = data.Id

	dataUpdates := make(map[string]interface{})
	dataUpdates["nama"] = data.Nama
	dataUpdates["password"] = data.Password
	dataUpdates["update_at"] = time.Now()

	err := r.queryrepo.UpdateData(&data, Where, dataUpdates)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) GetUserById(id string) (*[]models.Users, error) {
	var table []models.Users
	query := "SELECT * from users where id=" + id

	_, err := r.queryrepo.DinamicFindQueryRaw(&table, query)
	if err != nil {
		return nil, err
	}
	return &table, err

}
