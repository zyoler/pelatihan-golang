package main

import (
	"fmt"
	"net/http"
	"pertmuan_2/connection"
	"pertmuan_2/controller"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	koneksidb := connection.ConnectToDb()
	repo := repository.NewRepo(koneksidb)
	usecase := usecase.NewUsecase(repo)
	ctrl := controller.NewController(usecase)

	r.Group(func(rs chi.Router) {
		rs.Use(AuthorizationApikey, GetToken)
		rs.Get("/getdata", ctrl.GetData)
		rs.Post("/postdata", ctrl.PostData)
		rs.Delete("/delete-data-mhs/{id}", ctrl.DeleteData)
		rs.Put("/update-data-mhs/{id}", ctrl.UpdateData)

		rs.Get("/getdata-detail/{id}", ctrl.GetDataMhsDetail)
	})
	r.Group(func(rst chi.Router) {
		rst.Use(AuthorizationApikey)
		rst.Post("/login", GenerateTokens)
	})
	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5000", r); err != nil {
		fmt.Println("Error Starting Service")
	}
	fmt.Println("Starting Services")
}
