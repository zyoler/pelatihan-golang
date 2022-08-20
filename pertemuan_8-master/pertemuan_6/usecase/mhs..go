package usecase

import "pertemuan_6/models"

func (r *UC) GetDataMHS() ([]models.MhsPub, error) {
	var Modelmhs []models.MhsPub
	err := r.queryrepo.FindAll(Modelmhs)
	if err != nil {
		return Modelmhs, err
	}
	// var modelmhsd3 models.MhsD3
	// err = r.queryrepo.FindAll(modelmhsd3)
	// if err != nil {
	// 	return Modelmhs, err
	// }
	return Modelmhs, nil
}

func (r *UC) InsertDataMHS(data models.MhsPub) error {
	err := r.queryrepo.InsertData(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) UpdateDataMhs(data models.MhsPub) error {
	var Where map[string]interface{}
	Where["id"] = data.Id

	var dataUpdates map[string]interface{}
	dataUpdates["nama"] = data.Nama
	dataUpdates["umur"] = data.Umur

	err := r.queryrepo.UpdateData(&data, Where, dataUpdates)
	if err != nil {
		return err
	}
	return nil
}
func (r *UC) DeleteDataMhs(id int) error {
	var Where map[string]interface{}
	Where["id"] = id
	var TableMhs models.MhsPub
	err := r.queryrepo.DeleteData(&TableMhs, Where)
	if err != nil {
		return err
	}
	return nil
}
