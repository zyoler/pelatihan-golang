package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pertemuan/connection"
	"pertemuan/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB

var REDIS *redis.Client

func init() {
	DB = connection.ConnectToDb()
	REDIS = connection.ConnectToRedis() // Load connection
}

func GetData(w http.ResponseWriter, r *http.Request) {
	// GET
	var listmhsdata []models.MhsPub
	ctx := context.Background()

	_, err := REDIS.Get(ctx, "list_data_mhs_dani").Result() // Get data ke redis dengan key list_data_mhs

	/*
		hit 1 -> jika redis kosong maka get ke database
		hit ke 2 -> karena redis nya udah keisi maka get nya dari redis
	*/

	if err != nil {
		log.Println("Redisnya Kosong Get Database Dulu !!")
		// Kondisi saat redis kosong
		DB.Find(&listmhsdata)
		datajson, err := json.Marshal(listmhsdata)
		if err != nil {
			w.Write([]byte("Error Convert To JSON"))
			w.WriteHeader(500)
			return
		}

		err = REDIS.Set(ctx, "list_data_mhs_dani", (datajson), 0).Err()
		if err != nil {
			log.Println("Redis Error", err)
			log.Println("Error set ke Redis")
		}

		w.Header().Set("content-type", "application/json") // return type data
		w.Write(datajson)                                  // return data
		w.WriteHeader(200)
	}

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
	ctx := context.Background()

	REDIS.Del(ctx, "list_data_mhs_dani") // Recommended

	w.Write([]byte("Sukses GET DATA METHOD POST"))
	w.WriteHeader(200)
	return
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	idmhs := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.MhsPub
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	DB.Model(models.MhsPub{}).Where("id = ?", idmhs).Updates(datarequest)
	ctx := context.Background()

	// REDIS.FlushAll(ctx).Result() // Hapus semua data
	REDIS.Del(ctx, "list_data_mhs_dani") // Recommended

	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
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
	ctx := context.Background()

	REDIS.Del(ctx, "list_data_mhs_dani")
	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !!")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {
	// ex : 200 ok, 500 errpr / internal server error, 404 Bad Request
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
