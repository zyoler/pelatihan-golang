package controller

import (
	"Itemku/models"
	"encoding/json"
	"net/http"
)

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {
	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resevice)
}
