package app

import (
	"net/http"

	"github.com/crisguitar/api-golang/internal/app/product"
	"github.com/go-chi/chi"
)

func dummyHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("pong"))
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/dummy", dummyHandler)
	r.Get("/product", product.NewProductHandler)
	r.Post("/product", product.NewCreateProductHandler)
	return r
}
