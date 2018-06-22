package persistence

import "some-api/internal/app/product"

type Repository struct {
	products []product.Product
}

func (repository *Repository) List() []product.Product {
	return repository.products
}

func (repository *Repository) Save(newProduct product.Product) {
	repository.products = append(repository.products, newProduct)
}

func NewProductRepository() *Repository {
	return &Repository{}
}