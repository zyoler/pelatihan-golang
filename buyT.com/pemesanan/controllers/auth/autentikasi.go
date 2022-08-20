package auth

import (
	"encoding/json"
	"net/http"
	"pemesanan/config"
	"pemesanan/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.Connect()
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		Respon(w, 400, nil, "Bad Request")
		return
	}

	DB.Create(&user)
	Respon(w, 200, nil, "Register Success")
}

func Respon(w http.ResponseWriter, code int, data interface{}, message string) {
	Respon := models.Response{}
	Respon.Code = code
	Respon.Data = data
	Respon.Message = message
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Respon)
}
