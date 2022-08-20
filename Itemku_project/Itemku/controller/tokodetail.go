package controller

import (
	"Itemku/models"
	"Itemku/tools"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (c *ctrl) PostDataTokoDetail(w http.ResponseWriter, r *http.Request) {
	method := "POST"
	baseUrl := "http://localhost:8081/publish"

	decoder := json.NewDecoder(r.Body)
	var datarequest models.TokoDetail
	err := decoder.Decode(&datarequest)
	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}
	data := models.RequestNotify{
		IdPemesanan: datarequest.Id,
		Message:     "Success menambahkan toko detail",
		Data:        datarequest,
	}
	err = c.us.InsertDataTokoDetail(datarequest)
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

func (c *ctrl) DeleteDataTokoDetail(w http.ResponseWriter, r *http.Request) {
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
	err = c.us.DeleteDataTokoDetail(idinteger)
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

func (c *ctrl) GetDataTokoDetailById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	onlyOne, err := c.us.GetDataTokoDetailById(id)
	if err != nil {
		log.Println("isi err ", err)
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	ResponseApi(w, 200, onlyOne, "Sukses Get Data By Id")
}
