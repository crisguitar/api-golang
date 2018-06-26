package product

import (
	"net/http"
)

//go:generate mockgen -destination=../mocks/mock_handler.go -package=mocks github.com/crisguitar/api-golang/internal/app/product ResponseWriter
type ResponseWriter interface {
	http.ResponseWriter
}

func NewProductHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("[]"))
}

func NewCreateProductHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
