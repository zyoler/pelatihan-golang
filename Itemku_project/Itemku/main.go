package main

import (
	"Itemku/connection"
	"Itemku/controller"
	"Itemku/controller/auth"
	"Itemku/repository"
	"Itemku/usecase"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	koneksidb := connection.ConnectToDb()
	repo := repository.NewRepo(koneksidb)
	usecase := usecase.NewUsecase(repo)
	ctrl := controller.NewController(usecase)

	r.Group(func(rs chi.Router) {
		rs.Use(auth.GetToken)

		rs.Get("/getdata-game", ctrl.GetDataGame)

		rs.Get("/getdata-user", ctrl.GetDataUser)
		rs.Get("/getdata-user/{id}", ctrl.GetUserById)
		rs.Post("/postdata-user", ctrl.PostDataUser)
		rs.Delete("/deletedata-user/{id}", ctrl.DeleteDataUser)
		rs.Put("/updatedata-user/{id}", ctrl.UpdateDataUser)

		rs.Get("/getdata-toko", ctrl.GetDataToko)
		rs.Post("/postdata-toko", ctrl.PostDataToko)
		rs.Delete("/deletedata-toko/{id}", ctrl.DeleteDataToko)
		rs.Put("/updatedata-toko/{id}", ctrl.UpdateDataUser)

		rs.Post("/postdata-tokodetail", ctrl.PostDataTokoDetail)
		rs.Delete("/deletedata-tokodetail/{id}", ctrl.DeleteDataTokoDetail)

		rs.Get("/getdata-tokodetail/{id}", ctrl.GetDataTokoDetailById)

		rs.Get("/get-data-notif/{id}", ctrl.Notif)
	})
	r.Group(func(rst chi.Router) {
		rst.Use(auth.ApiKey)
		rst.Post("/login", auth.GenerateTokens)
	})
	fmt.Println("Running Service")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error Starting Service")
	}
	fmt.Println("Starting Services")
}
