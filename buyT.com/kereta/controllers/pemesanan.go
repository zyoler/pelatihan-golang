package controllers

import (
	"encoding/json"
	"kereta/config"
	"kereta/models"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.Connect()
}

func (c *BaseController) Pemesanan(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var pemesanan models.Pemesanan
	err := decoder.Decode(&pemesanan)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	DB.Create(&pemesanan)

	Respon(w, 200, pemesanan, "pemesanan Success")
}

func (c *BaseController) GetPemesanan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Pemesanan, err := c.us.FindAllPemesananById(id)
	if err != nil {
		Respon(w, 500, nil, "Internal Server Error")
		return
	}
	data, err := json.Marshal(Pemesanan)
	if err != nil {
		Respon(w, 500, nil, "Error marshall")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(200)
	return
}

// method := "POST"
// 	baseUrl := "http://localhost:8080/publish"
// 	err := json.NewDecoder(r.Body).Decode(&pemesanan)
// 	if err != nil {
// 		w.Write([]byte("Error Decode JSON Payload"))
// 		w.WriteHeader(500)
// 		return
// 	}

// 	data := models.RequestNotify{
// 		Message: "Selamat Pemesanan anda di proses",
// 		Data:    pemesanan,
// 	}

// 	dataString, _ := json.Marshal(data)
// 	code, result, err := tools.HTTPNotif(method, baseUrl, string(dataString), nil)
// 	if err != nil {
// 		Respon(w, code, nil, err.Error())
// 		return
// 	}
// 	log.Println(result)
