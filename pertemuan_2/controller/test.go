package controller

import (
	"encoding/json"
	"net/http"
	"pertemuan_2/models"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	// GET
	if r.Method == "GET" {
		data := GetMhs()
		// byte / json
		datajson, err := json.Marshal(data)

		if err != nil {
			w.Write([]byte("Error Convert To JSON"))
			w.WriteHeader(500)
			return
		}
		w.Header().Set("content-type", "application/json") // return type data
		w.Write(datajson)                                  // return data
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error not found", 404)
}

func PostData(w http.ResponseWriter, r *http.Request) {
	// POST
	if r.Method == "POST" {
		w.Write([]byte("Sukses GET DATA METHOD POST"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error not found", 404)
}

func GetMhs() []models.MhsPub {
	listMhsPasim := []models.MhsPub{
		{Id: 1, Nama: "Junad", Umur: 17},
		{Id: 2, Nama: "Junid", Umur: 18},
		{Id: 3, Nama: "Junud", Umur: 19},
	}
	return listMhsPasim
}
