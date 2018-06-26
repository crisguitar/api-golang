package product

import (
	"testing"
	"github.com/golang/mock/gomock"
	"net/http"
	. "github.com/onsi/gomega"
	"github.com/crisguitar/api-golang/internal/app/mocks"
)

func TestGetAllProducts(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Empty when there are no products", func(_ *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		mockResponseWriter := mocks.NewMockResponseWriter(mockController)

		header := http.Header(make(map[string][]string))

		mockResponseWriter.EXPECT().Header().Return(header).Times(1)
		mockResponseWriter.EXPECT().Write([]byte("[]")).Times(1)

		NewProductHandler(mockResponseWriter, nil)

		g.Expect(header["Content-Type"]).To(Equal([]string{"application/json"}))
	})
}