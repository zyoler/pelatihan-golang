package usecase

import (
	"Itemku/models"
)

func (r *UC) InsertDataTokoDetail(data models.TokoDetail) error {
	err := r.queryrepo.InsertData(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) DeleteDataTokoDetail(id int) error {
	Where := make(map[string]interface{})
	Where["id"] = id
	var Table models.TokoDetail
	err := r.queryrepo.DeleteData(&Table, Where)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) GetDataTokoDetailById(id string) (*[]models.TokoDetailItem, error) {
	var table []models.TokoDetailItem
	query := "SELECT a.*,b.* from tokos a inner join toko_details b on a.id=b.toko_id where a.id=" + id

	_, err := r.queryrepo.DinamicFindQueryRaw(&table, query)
	if err != nil {
		return nil, err
	}
	return &table, err
}
