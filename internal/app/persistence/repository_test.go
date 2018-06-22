package persistence

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"some-api/internal/app/product"
)

var _ = Describe("Repository", func() {

	Describe("when creating new repository", func() {
		It("should be empty", func() {

			repository := NewProductRepository()

			Expect(repository.List()).To(BeEmpty())
		})
	})

	Describe("when adding a new product", func() {
		It("should be persisted", func() {
			repository := NewProductRepository()

			repository.Save(product.Product{})
			Expect(repository.List()).To(Not(BeEmpty()))
		})
	})
})
