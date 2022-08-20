package main

import (
	"fmt"
	"net/http"
	"pemilu/controller"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/api/getdata/penyelenggara", controller.GetDataPenyelenggara)

	err := http.ListenAndServe(":5000", r)
	if err != nil {
		fmt.Println("Error starting service")
		return
	}

}
