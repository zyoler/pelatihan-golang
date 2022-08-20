package controller

import (
	"encoding/json"
	"net/http"
	"pertemuan/connection"
	"pertemuan/models"
	"strconv"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDb()
}

func GetData(w http.ResponseWriter, r *http.Request) {
	// GET
	var listmhsdata []models.MhsPub
	DB.Find(&listmhsdata)

	datajson, err := json.Marshal(listmhsdata)

	if err != nil {
		w.Write([]byte("Error Convert To JSON"))
		w.WriteHeader(500)
		return
	}
	w.Header().Set("content-type", "application/json") // return type data
	w.Write(datajson)                                  // return data
	w.WriteHeader(200)
}

func PostData(w http.ResponseWriter, r *http.Request) {
	// POST
	decoder := json.NewDecoder(r.Body)
	var datarequest models.MhsPub

	err := decoder.Decode(&datarequest)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	DB.Create(&datarequest)
	w.Write([]byte("Sukses GET DATA METHOD POST"))
	w.WriteHeader(200)
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	idmhs := chi.URLParam(r, "idmhs")
	var datarequest models.MhsPub
	if idmhs == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idinteger, err := strconv.Atoi(idmhs)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = idinteger
	DB.Delete(&datarequest)
	ResponseApi(w, 200, nil, "Sukses delete")
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {
	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resevice) //Mengconvert dari interface{} ke JSON
}

func GetMhs() []models.MhsPub {
	listMhsPasim := []models.MhsPub{
		{Id: 1, Nama: "Junad", Umur: 17},
		{Id: 2, Nama: "Junid", Umur: 18},
		{Id: 3, Nama: "Junud", Umur: 19},
	}
	return listMhsPasim
}
