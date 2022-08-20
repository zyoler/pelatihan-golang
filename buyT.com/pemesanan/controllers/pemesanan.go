package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pemesanan/config"
	"pemesanan/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

var REDIS *redis.Client

func init() {
	REDIS = config.RedisConnetion()
}

func (c *BaseController) GetDataKategori(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data, err := REDIS.Get(ctx, "Kategori").Result()
	if err != nil {
		log.Println("Get To Database")
		Kategori, err := c.uc.GetAllKategori()
		if err != nil {
			Respon(w, 500, nil, "Internal Server Error")
			return
		}
		data, err := json.Marshal(Kategori)
		if err != nil {
			Respon(w, 500, nil, "Error marshall")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		w.WriteHeader(200)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
	w.WriteHeader(200)
	return
}

func (c *BaseController) GetDataKereta(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data, err := REDIS.Get(ctx, "Kereta").Result()
	if err != nil {
		log.Println("Get To Database")
		Kereta, err := c.uc.GetAllKereta()
		if err != nil {
			Respon(w, 500, nil, "Internal Server Error")
			return
		}
		data, err := json.Marshal(Kereta)
		if err != nil {
			Respon(w, 500, nil, "Error marshall")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		w.WriteHeader(200)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
	w.WriteHeader(200)
	return
}

func (c *BaseController) GetDataDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data, err := REDIS.Get(ctx, "Detail").Result()
	if err != nil {
		log.Println("Get To Database")
		Detail, err := c.uc.GetAllDetail()
		if err != nil {
			Respon(w, 500, nil, "Internal Server Error")
			return
		}
		data, err := json.Marshal(Detail)
		if err != nil {
			Respon(w, 500, nil, "Error marshall")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		w.WriteHeader(200)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
	w.WriteHeader(200)
	return
}

func (c *BaseController) GetDataDetailById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		Respon(w, 400, nil, "Bad Request")
		return
	}
	id_kereta, err := strconv.Atoi(id)
	if err != nil {
		Respon(w, 400, nil, "Bad Request")
		return
	}
	data, err := c.uc.GetDetailById(id_kereta)
	if err != nil {
		Respon(w, 500, nil, "Internal Server Error")
		return
	}
	Respon(w, 200, data, "Success Get data Kereta")
}

func Respon(w http.ResponseWriter, code int, data interface{}, message string) {
	Respon := models.Response{}
	Respon.Code = code
	Respon.Data = data
	Respon.Message = message
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Respon)
}
