package repository

func (r *repo) FindAll(i interface{}) error {
	result := r.apps.Find(i)
	// DB.Find(&listmhsdata)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
