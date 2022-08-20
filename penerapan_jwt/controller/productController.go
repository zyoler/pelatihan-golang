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

var DBKU *gorm.DB

var REDISPRODUCT *redis.Client

func init() {
	DBKU = connection.ConnectToDb()
	REDISPRODUCT = connection.ConnectToRedis()
}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	ctx := context.Background()
	redisData, err := REDISPRODUCT.Get(ctx, "listproducts").Result()
	if err != nil {
		log.Println("Redisnya Kosong Get Database Dulu  !! ")
		DBKU.Find(&products)
		log.Println("isi ", products)
		datajson, err := json.Marshal(products)
		log.Println("ss", err)
		if err != nil {
			log.Println("Error Convert Json", err)
			w.Write([]byte("Error Convert TO JSON"))
			w.WriteHeader(500)
			return
		}
		err = REDISPRODUCT.Set(ctx, "listproducts", (datajson), 0).Err()
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

func PostProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Product
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}
	DBKU.Create(&datarequest)
	ctx := context.Background()
	REDISPRODUCT.Del(ctx, "listproducts") // Recomended
	ResponseApi(w, 500, nil, "Sukses Insert Data")
	return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := chi.URLParam(r, "id")
	var datarequest models.Product
	if idProduct == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idinteger, err := strconv.Atoi(idProduct)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = idinteger
	DB.Delete(&datarequest)
	ctx := context.Background()
	REDISPRODUCT.Del(ctx, "listproducts") // Recomended

	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Product
	err := decoder.Decode(&datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	DB.Model(models.Product{}).Where("id = ?", idProduct).Updates(datarequest)
	ctx := context.Background()
	REDISPRODUCT.Del(ctx, "listproducts")
	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}
