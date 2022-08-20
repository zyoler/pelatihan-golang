package usecase

import (
	"log"
	"notification/models"
)

func (r *UC) NotifMhs(id int) (models.Notif, error) {
	var notifObj models.Notif
	where := map[string]interface{}{}
	where["mhs_id"] = id
	err := r.queryrepo.FindOne(&notifObj, where)
	if err != nil {
		log.Println("Erro query ", err)
		return notifObj, err
	}
	return notifObj, nil
}
