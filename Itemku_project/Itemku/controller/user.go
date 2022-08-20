package controller

import (
	"Itemku/connection"
	"Itemku/models"
	"Itemku/tools"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

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

func (c *ctrl) GetDataUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redisData, err := REDIS.Get(ctx, "user").Result()
	if err != nil {
		listdata, err := c.us.GetDataUser()
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
		}
		datajson, err := json.Marshal(listdata)
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
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

func (c *ctrl) PostDataUser(w http.ResponseWriter, r *http.Request) {
	method := "POST"
	baseUrl := "http://localhost:8081/publish"

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Users
	err := decoder.Decode(&datarequest)
	datarequest.CreateAt = time.Now()
	datarequest.UpdateAt = time.Now()
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}
	data := models.RequestNotify{
		IdPemesanan: datarequest.Id,
		Message:     "success menambahkan user",
		Data:        datarequest,
	}
	err = c.us.InsertDataUser(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, err.Error())
		return
	}
	dataString, _ := json.Marshal(data)
	code, result, err := tools.HTTPNotif(method, baseUrl, string(dataString), nil)
	if err != nil {
		ResponseApi(w, code, nil, err.Error())
		return
	}
	log.Println(result)
	ctx := context.Background()
	REDIS.Del(ctx, "user")
	ResponseApi(w, 200, nil, "Sukses Insert Data")
}

func (c *ctrl) DeleteDataUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	idinteger, err := strconv.Atoi(id)
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
	REDIS.Del(ctx, "user")
	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
}

func (c *ctrl) UpdateDataUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Users
	err := decoder.Decode(&datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}
	idinteger, err := strconv.Atoi(id)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	datarequest.Id = idinteger
	err = c.us.UpdateDataUser(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	ctx := context.Background()
	REDIS.Del(ctx, "user")
	ResponseApi(w, 200, nil, "Sukses Update Data")
}

func (c *ctrl) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	onlyOne, err := c.us.GetUserById(id)
	if err != nil {
		log.Println("isi err ", err)
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ResponseApi(w, 200, onlyOne, "Sukses Get Data By Id")
}
