package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"penerepan_jwt/connection"
	"penerepan_jwt/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB

var REDIS *redis.Client

func init() {
	DB = connection.ConnectToDb()
	REDIS = connection.ConnectToRedis()
}
func GetKatagoris(w http.ResponseWriter, r *http.Request) {
	var kategoris []models.Kategori
	ctx := context.Background()

	redisData, err := REDIS.Get(ctx, "listKategoris").Result()
	if err != nil {
		log.Println("Redisnya Kosong Get Database Dulu  !! ")
		DB.Find(&kategoris)
		log.Println("isi ", kategoris)
		datajson, err := json.Marshal(kategoris)
		log.Println("ss", err)
		if err != nil {
			log.Println("Error Convert Json", err)
			w.Write([]byte("Error Convert TO JSON"))
			w.WriteHeader(500)
			return
		}
		err = REDIS.Set(ctx, "listKategoris", (datajson), 0).Err()
		if err != nil {
			log.Println("Redis Error", err)
			log.Println("Error Set Ke Redis")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}

	log.Println("Redisnya ada Ga get Database tapi  Get Ke Redis   !! ")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(redisData))
	w.WriteHeader(200)
}

func PostKategoris(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Kategori
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}
	DB.Create(&datarequest)
	ctx := context.Background()
	REDIS.Del(ctx, "listKategoris")

	ResponseApi(w, 500, nil, "Sukses Insert Data")
	return
}

func DeleteKategoris(w http.ResponseWriter, r *http.Request) {
	idKtg := chi.URLParam(r, "id")
	var datarequest models.Kategori

	if idKtg == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idinteger, err := strconv.Atoi(idKtg)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = idinteger
	DB.Delete(&datarequest)
	ctx := context.Background()
	REDIS.Del(ctx, "listKategoris")

	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func UpdateKategoris(w http.ResponseWriter, r *http.Request) {
	idKtg := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Kategori
	err := decoder.Decode(&datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	DB.Model(models.Kategori{}).Where("id = ?", idKtg).Updates(datarequest)
	ctx := context.Background()
	REDIS.Del(ctx, "listKategoris")
	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {
	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resevice)
}
