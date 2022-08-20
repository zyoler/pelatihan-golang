package controller

import (
	"Itemku/models"
	"Itemku/tools"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (c *ctrl) GetDataToko(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redisData, err := REDIS.Get(ctx, "toko").Result()
	if err != nil {
		listdata, err := c.us.GetDataToko()
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

func (c *ctrl) PostDataToko(w http.ResponseWriter, r *http.Request) {
	method := "POST"
	baseUrl := "http://localhost:8081/publish"

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Toko
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
		Message:     "Success menambahkan toko",
		Data:        datarequest,
	}
	err = c.us.InsertDataToko(datarequest)
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
	REDIS.Del(ctx, "toko")
	ResponseApi(w, 200, nil, "Sukses Insert Data")
}

func (c *ctrl) DeleteDataToko(w http.ResponseWriter, r *http.Request) {
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
	err = c.us.DeleteDataToko(idinteger)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ctx := context.Background()
	REDIS.Del(ctx, "toko")
	if err != nil {
		ResponseApi(w, 500, nil, "Error Delete !! ")
		return
	}
	ResponseApi(w, 200, nil, "Sukses delete")
}

func (c *ctrl) UpdateDataToko(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Toko
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
	err = c.us.UpdateDataToko(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}
	ctx := context.Background()
	REDIS.Del(ctx, "toko")
	ResponseApi(w, 200, nil, "Sukses Update Data")
}
