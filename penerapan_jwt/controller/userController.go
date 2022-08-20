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

var DBUSER *gorm.DB

var REDISUSER *redis.Client

func init() {
	DBUSER = connection.ConnectToDb()
	REDISUSER = connection.ConnectToRedis()
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	ctx := context.Background()

	redisData, err := REDIS.Get(ctx, "listusers").Result()
	if err != nil {
		log.Println("Redisnya Kosong Get Database Dulu  !! ")
		DBUSER.Find(&users)
		log.Println("isi ", users)
		datajson, err := json.Marshal(users)
		log.Println("ss", err)
		if err != nil {
			log.Println("Error Convert Json", err)
			w.Write([]byte("Error Convert TO JSON"))
			w.WriteHeader(500)
			return
		}
		err = REDIS.Set(ctx, "listusers", (datajson), 0).Err()
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

func PostUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.User
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}
	DBUSER.Create(&datarequest)
	ctx := context.Background()
	REDISUSER.Del(ctx, "listusers")

	ResponseApi(w, 500, nil, "Sukses Insert Data")
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "id")
	var datarequest models.User

	if idUser == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idinteger, err := strconv.Atoi(idUser)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = idinteger
	DBUSER.Delete(&datarequest)
	ctx := context.Background()
	REDISUSER.Del(ctx, "listusers")

	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.User
	err := decoder.Decode(&datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	DBUSER.Model(models.User{}).Where("id = ?", idUser).Updates(datarequest)
	ctx := context.Background()
	REDISUSER.Del(ctx, "listusers")
	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}
