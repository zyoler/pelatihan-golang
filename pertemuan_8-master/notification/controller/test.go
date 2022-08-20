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
	// var listmhsdata []models.MhsPub
	ctx := context.Background()

	redisData, err := REDIS.Get(ctx, "list_data_mhs").Result() // Get Data Ke Redis dengan key list_data_mhs

	/*
		hit 1  -> jika redis kosong maka get ke database
		hit ke 2  -> karna redis nya udah keisi maka get nya dari redis
	*/
	if err != nil {
		log.Println("Redisnya Kosong Get Database Dulu  !! ")
		// kondisi Saat Redis Kosong
		listmhsdata, err := c.us.GetDataMHS()
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
		}
		datajson, err := json.Marshal(listmhsdata)
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
		}

		err = REDIS.Set(ctx, "list_data_mhs", (datajson), 0).Err()
		if err != nil {
			log.Println("Redis Error", err)
			log.Println("Error Set Ke Redis")
		}
		w.Header().Set("Content-Type", "application/json") //return type data nya
		w.Write(datajson)                                  //return datanya
		w.WriteHeader(200)
		return

	}

	log.Println("Redisnya ada Ga get Database tapi  Get Ke Redis   !! ")
	w.Header().Set("Content-Type", "application/json") //return type data nya
	w.Write([]byte(redisData))                         //return datanya
	w.WriteHeader(200)
	// return
	// ResponseApi(w, 200, []byte(redisData), "Sukses Get Data")
}

func (c *ctrl) PostData(w http.ResponseWriter, r *http.Request) {
	// post / ge
	decoder := json.NewDecoder(r.Body)
	var datarequest models.MhsPub
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	err = c.us.InsertDataMHS(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ctx := context.Background()

	REDIS.Del(ctx, "list_data_mhs")

	ResponseApi(w, 200, nil, "Sukses Insert Data")
	return
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

	err = c.us.DeleteDataMhs(idinteger)

	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ctx := context.Background()
	REDIS.Del(ctx, "list_data_mhs") // Recomended
	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func (c *ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	idmhs := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.MhsPub
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
	datarequest.Id = idinteger
	err = c.us.UpdateDataMhs(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	ctx := context.Background()

	REDIS.Del(ctx, "list_data_mhs") // Recomended

	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {
	// nerima semua bentuk models -> models mhs , models warga , models buku,nil
	// ex : 200 -> ok ,500 -> error / internal server ,400 bad request  ,401 -> unauthorized , 404 -> data not found
	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resevice)
}
