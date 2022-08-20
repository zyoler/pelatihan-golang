package repository

import "database/sql"

func (r *repo) FindAll(i interface{}) error {
	result := r.apps.Find(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) InsertData(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) UpdateData(i interface{}, where map[string]interface{}, data map[string]interface{}) error {
	result := r.apps.Model(i).Where(where).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) DeleteData(i interface{}, where map[string]interface{}) error {
	result := r.apps.Where(where).Delete(i)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) DinamicFindQueryRaw(i interface{}, query string) (*sql.Rows, error) {
	rows, err := r.apps.Raw(query).Rows()
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		r.apps.ScanRows(rows, i)
	}

	return rows, nil
}
