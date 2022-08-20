package main

import (
	"fmt"
	"net/http"
	"pertemuan/controller"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/getdata", controller.GetData)
	r.Post("/postdata", controller.PostData)
	r.Get("/delete-data-mhs/{idmhs}", controller.DeleteData)

	fmt.Println("Running Service")
	if err := http.ListenAndServe(":5000", r); err != nil {
		fmt.Println("Error Starting Service")
		return
	}
	fmt.Println("Starting Services")
}
