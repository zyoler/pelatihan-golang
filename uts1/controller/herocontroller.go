package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"uts_1/connection"
	"uts_1/models"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB

var REDIS *redis.Client

func init() {
	DB = connection.ConnectDb()
	REDIS = connection.ConnectRedis()
}

func GetHero(w http.ResponseWriter, r *http.Request) {
	var listhero []models.HeroML
	ctx := context.Background()

	redisData, err := REDIS.Get(ctx, "list_hero").Result()

	if err != nil {
		log.Println("Redis is empty, create a new redis")
		DB.Find(&listhero)
		datajson, err := json.Marshal(listhero)
		if err != nil {
			log.Println("Error convert to JSON", err)
			w.Write([]byte("Error convert to JSON"))
			w.WriteHeader(500)
			return
		}

		err = REDIS.Set(ctx, "list_hero", (datajson), 0).Err()
		if err != nil {
			log.Println("Redis Error", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(redisData))
	w.WriteHeader(200)
}

func PostHero(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.HeroML
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}
	DB.Create(&datarequest)
	ctx := context.Background()
	REDIS.Del(ctx, "list_hero")

	ResponseApi(w, 500, nil, "Success Insert Methode")
}

func DeleteHero(w http.ResponseWriter, r *http.Request) {
	idhero := chi.URLParam(r, "id")
	var datarequest models.HeroML

	if idhero == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idint, err := strconv.Atoi(idhero)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = idint
	DB.Delete(&datarequest)
	ctx := context.Background()

	REDIS.Del(ctx, "list_hero")
	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Success Delete Methode")
}

func UpdateHero(w http.ResponseWriter, r *http.Request) {
	idhero := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.HeroML
	err := decoder.Decode(&datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	DB.Model(models.HeroML{}).Where("id = ?", idhero).Updates(datarequest)
	ctx := context.Background()

	REDIS.Del(ctx, "list_hero")
	ResponseApi(w, 200, nil, "Success Update Methode")
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
