package app

import 	"github.com/go-chi/chi"
import (
	"net/http"
	"github.com/crisguitar/api-golang/internal/app/product"
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {

}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/dummy", dummyHandler)
	r.Get("/product", product.NewProductHandler)
	r.Post("/product", product.NewCreateProductHandler)
	return r
}
