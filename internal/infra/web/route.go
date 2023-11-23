package web

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/health"))

	//mux.Use(app.SessionLoad)

	mux.Post("/product", app.CreateProductHandler)
	mux.Get("/product", app.ListProductsHandler)
	mux.Get("/product/{id}", app.GetProductsHandler)

	mux.Post("/client", app.CreateClientHandler)
	mux.Get("/client/{id}", app.GetClientHandler)

	mux.Post("/insert-item", app.AddToCart)
	mux.Get("/checkout/{id}", app.CheckoutHandler)
	mux.Post("/payment", app.PaymentHandler)

	return mux
}
