package response

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/models/responseModel"
)

func Response(w http.ResponseWriter, code int, msg string, data interface{}) {
	receiver := responseModel.Response{}

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

	w.WriteHeader(code) // status code
	w.Write(jadiJson)   // return datanya
}

func cekStatus(code int) string {
	if code == 200 {
		return "SUKSES"
	}
	return "GAGAL"
}

func MsgTambah(model string) string {
	return "Berhasil menambahkan data " + model
}

func MsgGetAll(model string) string {
	return "Berhasil mengambil semua data " + model
}

func MsgGetDetail(model string) string {
	return "Berhasil mengambil detail data " + model
}

func MsgUpdate(model string) string {
	return "Berhasil mengupdate data " + model
}

func MsgHapus(model string) string {
	return "Berhasil menghapus data " + model
}

func MsgNotFound(model string) string {
	return "Data " + model + " tidak ditemukan"
}

func MsgServiceErr() string {
	return "Terdapat masalah pada service"
}

func MsgInvalidReq() string {
	return "Invalid Request"
}
