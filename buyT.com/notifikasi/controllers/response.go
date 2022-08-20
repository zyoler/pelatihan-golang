package controllers

import (
	"encoding/json"
	"net/http"
	"notifikasi/models"
)

var (
	success               = "SUCCESS"
	bad_request           = "BAD REQUEST"
	Internal_server_error = "INTERNAL SERVER ERROR"
	not_found             = "NOT FOUND"
	unknown               = "UNKNOWN"
	service_running       = "SERVICE RUNNING"
)

func cekStatus(code int) string {
	if code == http.StatusOK {
		return success
	} else if code == http.StatusBadRequest {
		return bad_request
	} else if code == http.StatusInternalServerError {
		return Internal_server_error
	} else if code == http.StatusNotFound {
		return not_found
	} else {
		return unknown
	}
}

func MsgGetAll(exist bool, model string) string {
	if exist {
		return "Berhasil mengambil semua data " + model
	}
	return "Gagal mengambil semua data " + model
}

func MsgGetDetail(exist bool, model string) string {
	if exist {
		return "Berhasil mengambil detail data " + model
	}
	return "Detail data " + model + " tidak ditemukan"
}

func MsgCreate(exist bool, model string) string {
	if exist {
		return "Berhasil menambahkan data " + model
	}
	return "Gagal menambahkan data " + model
}

func MsgUpdate(exist bool, model string) string {
	if exist {
		return "Berhasil mengupdate data " + model
	}
	return "Gagal mengupdate data " + model
}

func MsgDelete(exist bool, model string) string {
	if exist {
		return "Berhasil menghapus data " + model
	}
	return "Gagal menghapus data " + model
}

func Response(w http.ResponseWriter, code int, msg string, data interface{}) {
	receiver := models.Response{}

	receiver.Status = cekStatus(code)
	receiver.Code = code
	receiver.Message = msg
	receiver.Data = data

	jadiJson, err := json.Marshal(receiver) // nge convert jadi json

	w.Header().Set("Content-Type", "application/json") // return type data nya

	if err != nil {
		w.WriteHeader(500) // status code
		w.Write([]byte("Error to marshall"))
	}

	// success
	w.WriteHeader(code) // status code
	w.Write(jadiJson)   // return datanya
}

func ResponseRunningService(w http.ResponseWriter) {
	Response(w, http.StatusOK, service_running, nil)
}
