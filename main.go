package main

import (
	"net/http"
	"pb2/rest"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/getmyvehicle", rest.GetMyVehicle)
		r.Post("/postmyvehicle", rest.PostMyVehicle)
	})
	http.ListenAndServe(":8080", router)
}
