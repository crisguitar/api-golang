package persistence

import (
	"github.com/icrowley/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/api-golang/internal/app/product"
)

var _ = Describe("Repository", func() {

	var repository *Repository

	BeforeEach(func() {
		repository = NewProductRepository()
	})

	Describe("when creating new repository", func() {
		It("should be empty", func() {
			Expect(repository.List()).To(BeEmpty())
		})
	})

	Describe("when adding a new product", func() {
		It("should be persisted", func() {
			product := &Product{Colour:fake.Color(), FabricType: fake.Word(), Season: fake.Word(), Status: fake.Word()}
			savedProduct := repository.Save(product)
			Expect(repository.List()).To(Not(BeEmpty()))
			Expect(savedProduct.Colour).To(Equal(product.Colour))
			Expect(savedProduct.FabricType).To(Equal(product.FabricType))
			Expect(savedProduct.Season).To(Equal(product.Season))
			Expect(savedProduct.Status).To(Equal(product.Status))
		})

		It("should have an Id", func() {
			newProduct := repository.Save(&Product{})
			Expect(newProduct.Id).To(Not(BeNil()))
		})

		It("can be retrieved by Id", func() {
			savedProduct := repository.Save(&Product{})
			retrievedProduct := repository.Get(savedProduct.Id)
			Expect(retrievedProduct).To(Equal(savedProduct))
		})

		It("saves all fields", func() {
			savedProduct := repository.Save(&Product{Colour: fake.Color(), FabricType: fake.Word(), Season: fake.Word(), Status: fake.Month()})
			retrievedProduct := repository.Get(savedProduct.Id)
			Expect(retrievedProduct.Colour).To(Equal(savedProduct.Colour))
			Expect(retrievedProduct.FabricType).To(Equal(savedProduct.FabricType))
			Expect(retrievedProduct.Season).To(Equal(savedProduct.Season))
			Expect(retrievedProduct.Status).To(Equal(savedProduct.Status))
		})
	})
})
