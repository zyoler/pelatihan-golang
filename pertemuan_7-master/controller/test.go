package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pertemuan_6/connection"
	"pertemuan_6/models"
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
		log.Println("isi ", listmhsdata)
		datajson, err := json.Marshal(listmhsdata)
		log.Println("ss", err)
		if err != nil {
			log.Println("Error Convert Json", err)
			w.Write([]byte("Error Convert TO JSON"))
			w.WriteHeader(500)
			return
		}
		// log.Println("isina", datajson)
		//
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
	DB.Create(&datarequest)
	ctx := context.Background()

	// REDIS.FlushAll(ctx).Result() // hapus Semua Data (cara yang ga recomended)
	REDIS.Del(ctx, "list_data_mhs") // Recomended

	ResponseApi(w, 200, nil, "Sukses Insert Data")
	return
}

func (c *ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	idmhs := chi.URLParam(r, "id")
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

	// REDIS.FlushAll(ctx).Result()    // hapus Semua Data
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

	DB.Model(models.MhsPub{}).Where("id = ?", idmhs).Updates(datarequest)
	ctx := context.Background()

	// REDIS.FlushAll(ctx).Result()    // hapus Semua Data
	REDIS.Del(ctx, "list_data_mhs") // Recomended
	// " NULL
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
func GetMhs() []models.MhsPub {
	listMhsPasim := []models.MhsPub{
		{Id: 1, Nama: "Juned", Umur: 17},
		{Id: 2, Nama: "Frendi", Umur: 19},
		{Id: 3, Nama: "Fraky", Umur: 20},
	}

	return listMhsPasim
}
