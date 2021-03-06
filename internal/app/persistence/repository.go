package persistence

import (
	"github.com/crisguitar/api-golang/internal/app/product"
	"github.com/satori/go.uuid"
)

type Repository struct {
	products map[uuid.UUID]*product.Product
}

func (repository *Repository) List() map[uuid.UUID]*product.Product {
	return repository.products
}

func (repository *Repository) Save(newProduct *product.Product) *product.Product {
	newProduct.Id = uuid.NewV4()
	repository.products[newProduct.Id] = newProduct
	return newProduct
}

func (repository *Repository) Get(productId uuid.UUID) *product.Product {
	return repository.products[productId]
}
func (repository *Repository) Delete(uuid uuid.UUID) {
	delete(repository.products, uuid)
}

func NewProductRepository() *Repository {
	return &Repository{products: make(map[uuid.UUID]*product.Product)}
}
