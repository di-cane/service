package main

import (
	"dicane-api/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// Specify who is allowed
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/sales", controller.GetAllSales)
	mux.Get("/sales/{id}", controller.GetSaleById)
	mux.Post("/sales", controller.InsertSale)
	mux.Put("/sales/{id}", controller.UpdateSale)
	mux.Delete("/sales/{id}", controller.DeleteSale)

	// Priority routes
	mux.Get("/priority/list/{sale_id}", controller.GetPriorityList)
	// mux.Post("/priority/list", controller.InsertPriorityList)

	// Breeder routes
	mux.Post("/breeder", controller.InsertBreeder)
	mux.Get("/breeder/{email}", controller.GetBreederByEmail)
	mux.Get("/breeder/{id}/sales", controller.GetBreederSales)

	return mux
}
