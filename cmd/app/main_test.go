package app

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/crisguitar/api-golang/internal/app/product"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router Integration Test", func() {

	server := httptest.NewServer(NewRouter())

	AfterSuite(func() {
		server.Close()
	})

	It("Should have status code 200 for dummy endpoint", func() {
		response, _ := http.Get(server.URL + "/dummy")
		Expect(response.StatusCode).To(Equal(200))
		body, _ := ioutil.ReadAll(response.Body)
		Expect(string(body)).To(Equal("pong"))
	})

	Describe("Product Integration Test", func() {
		It("Should return empty list of products", func() {
			response, _ := http.Get(server.URL + "/product")
			Expect(response.StatusCode).To(Equal(200))
			body, _ := ioutil.ReadAll(response.Body)
			Expect(string(body)).To(Equal("[]"))
		})

		It("should save one product", func() {
			product := &product.Product{}
			body, _ := json.Marshal(product)
			requestBody := bytes.NewReader(body)
			response, _ := http.Post(server.URL+"/product", "application/json", requestBody)
			responseBody, _ := ioutil.ReadAll(response.Body)
			Expect(response.StatusCode).To(Equal(201))
			Expect(string(responseBody)).To(Equal(""))
		})
	})
})
