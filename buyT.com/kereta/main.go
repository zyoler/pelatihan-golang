package main

import (
	"kereta/config"
	"kereta/controllers"
	"kereta/controllers/auth"
	"kereta/repository"
	"kereta/usecase"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	connection := config.Connect()
	repository := repository.NewRepository(connection)
	usecase := usecase.NewUsecase(repository)
	controller := controllers.NewController(usecase)

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(auth.GetTokenJwt)
		r.Get("/api/s1/stasiuns", controller.GetDataStasiun)
		r.Get("/api/s1/stasiuns/{id}", controller.GetStasiunById)
		r.Get("/api/k1/kota/{id}", controller.GetKotaById)
		r.Get("/api/k2/kereta/pemesanan/{id}", controller.GetServisPemesanan)
		r.Post("/api/p1/kereta/pemesanan", controller.Pemesanan)
		r.Get("/api/p2/kereta/pemesanan/{id}", controller.GetPemesanan)
	})

	router.Group(func(r chi.Router) {
		r.Use(auth.ApiKey)
		r.Post("/api/au1/register", auth.Register)
		r.Post("/api/au2/login", auth.Login)
	})

	if err := http.ListenAndServe(":8001", router); err != nil {
		log.Fatal(err)
	}
	log.Println("Server Running on port 8001")
}
