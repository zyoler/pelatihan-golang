package controllers

import (
	"context"
	"encoding/json"
	"kereta/config"
	"kereta/models"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

var REDIS *redis.Client

func init() {
	REDIS = config.RedisConnetion()
}

func (c *BaseController) GetDataStasiun(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data, err := REDIS.Get(ctx, "stasiun").Result()
	if err != nil {
		log.Println("Get To Database")
		Stasiun, err := c.us.Stasiun()
		if err != nil {
			Respon(w, 500, nil, "Internal Server Error")
			return
		}
		data, err := json.Marshal(Stasiun)
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

func (c *BaseController) GetStasiunById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Stasiun, err := c.us.GetFirstStasiun(id)
	if err != nil {
		Respon(w, 500, nil, "Internal Server Error")
		return
	}
	data, err := json.Marshal(Stasiun)
	if err != nil {
		Respon(w, 500, nil, "Error marshall")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(200)
	return
}

func (c *BaseController) GetKotaById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Kota, err := c.us.GetAllStasiunByKota(id)
	if err != nil {
		Respon(w, 500, nil, "Internal Server Error")
		return
	}
	data, err := json.Marshal(Kota)
	if err != nil {
		Respon(w, 500, nil, "Error marshall")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(200)
	return
}

// func (c *BaseController) Notif(w http.ResponseWriter, r *http.Request) {
// 	method := "GET"
// 	baseUrl := "http://localhost:8000/notifikasi-stasiun/n1/1/1"
// 	header := map[string][]string{
// 		"Content-type": {"application/json"},
// 	}
// 	var ResponseServiceNotif models.ResponseNotifService

// 	err := tools.Curl(method, baseUrl, nil, &ResponseServiceNotif, header)
// 	if err != nil {
// 		log.Println("Error Hit Service Notification=> ", err)
// 		Respon(w, 500, nil, "Internal Server Error")
// 		return
// 	}

// 	Respon(w, 200, ResponseServiceNotif.Notif, "Get Data Success")

// }

func Respon(w http.ResponseWriter, code int, data interface{}, message string) {
	Respon := models.Response{}
	Respon.Code = code
	Respon.Data = data
	Respon.Message = message
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Respon)
}
