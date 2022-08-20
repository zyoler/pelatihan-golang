package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"notification/connection"
	"notification/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB

var REDIS *redis.Client

func init() {
	DB = connection.ConnectToDb()
	REDIS = connection.ConnectToRedis() // Load Connection
}

func (c *ctrl) GetData(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	redisData, err := REDIS.Get(ctx, "list_user").Result()

	if err != nil {
		listuser, err := c.us.GetDataUser()
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
		}
		datajson, err := json.Marshal(listuser)
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
		}

		err = REDIS.Set(ctx, "list_user", (datajson), 0).Err()
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

func (c *ctrl) PostData(w http.ResponseWriter, r *http.Request) {
	// post / ge
	decoder := json.NewDecoder(r.Body)
	var datarequest models.User
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	err = c.us.InsertDataUser(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ctx := context.Background()

	REDIS.Del(ctx, "list_user")

	ResponseApi(w, 200, nil, "Sukses Insert Data")
}

func (c *ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	idmhs := chi.URLParam(r, "id")
	if idmhs == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idinteger, err := strconv.Atoi(idmhs)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	err = c.us.DeleteDataUser(idinteger)

	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ctx := context.Background()
	REDIS.Del(ctx, "list_user") // Recomended
	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
}

func (c *ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	idmhs := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.User
	err := decoder.Decode(&datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}
	idinteger, err := strconv.Atoi(idmhs)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.IdUser = idinteger
	err = c.us.UpdateDataUser(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	ctx := context.Background()

	REDIS.Del(ctx, "list_user") // Recomended

	ResponseApi(w, 200, nil, "Sukses Update Data")
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {
	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resevice)
}
