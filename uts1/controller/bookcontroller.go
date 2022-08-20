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
)

func init() {
	DB = connection.ConnectDb()
	REDIS = connection.ConnectRedis()
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var listbook []models.Books
	ctx := context.Background()

	redisData, err := REDIS.Get(ctx, "list_book").Result()

	if err != nil {
		log.Println("Redis is empty, create a new redis")
		DB.Find(&listbook)
		datajson, err := json.Marshal(listbook)
		if err != nil {
			log.Println("Error convert to JSON", err)
			w.Write([]byte("Error convert to JSON"))
			w.WriteHeader(500)
			return
		}

		err = REDIS.Set(ctx, "list_book", (datajson), 0).Err()
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

func PostBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Books
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}
	DB.Create(&datarequest)
	ctx := context.Background()
	REDIS.Del(ctx, "list_book")

	ResponseApi(w, 500, nil, "Success Insert Methode")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idbook := chi.URLParam(r, "id")
	var datarequest models.Books

	if idbook == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idint, err := strconv.Atoi(idbook)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = idint
	DB.Delete(&datarequest)
	ctx := context.Background()

	REDIS.Del(ctx, "list_book")
	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Success Delete Methode")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idbook := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Books
	err := decoder.Decode(&datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	DB.Model(models.Books{}).Where("id = ?", idbook).Updates(datarequest)
	ctx := context.Background()

	REDIS.Del(ctx, "list_book")
	ResponseApi(w, 200, nil, "Success Update Methode")
}
