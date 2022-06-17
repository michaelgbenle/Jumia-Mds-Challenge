package test

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/michaelgbenle/jumiaMds/config"
	mockdatabase "github.com/michaelgbenle/jumiaMds/database/mocks"
	"github.com/michaelgbenle/jumiaMds/handlers"
	"github.com/michaelgbenle/jumiaMds/models"
	"github.com/michaelgbenle/jumiaMds/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	config.NewConfig("../../.env")
	os.Exit(m.Run())
}

func TestGetProductBySku(t *testing.T) {
	ctrl := gomock.NewController(t)
	//creates a new mock instance
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handlers.Handler{DB: mockDB}

	route := router.SetupRouter(h)
	product := models.Product{
		Name:    "Samsung Phone",
		Sku:     "cbf87a9be799",
		Stock:   2,
		Country: "ma",
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for error", func(t *testing.T) {
		mockDB.EXPECT().GetProductSku(product.Sku, product.Country).Return(nil, errors.New("error exist"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/product?sku=cbf87a9be799&country=ma", strings.NewReader(string(productJSON)))
		route.ServeHTTP(w, req)
		assert.Contains(t, w.Body.String(), "error fetching data")
		assert.Equal(t, w.Code, 500)

	})

	t.Run("Testing for successful request", func(t *testing.T) {
		mockDB.EXPECT().GetProductSku(product.Sku, product.Country).Return(&product, nil).Times(1)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/product?sku=cbf87a9be799&country=ma", strings.NewReader(string(productJSON)))
		route.ServeHTTP(w, req)
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Contains(t, w.Body.String(), string(productJSON))
	})

}

func TestConsumeStock(t *testing.T) {
	ctrl := gomock.NewController(t)
	//creates a new mock instance
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handlers.Handler{DB: mockDB}

	route := router.SetupRouter(h)
	product := models.Product{
		Name:    "Samsung Phone",
		Sku:     "cbf87a9be799",
		Stock:   10,
		Country: "ma",
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		t.Fail()
	}
	order := models.Order{
		ProductId: 1,
		Quantity:  10,
	}
	orderJSON, err := json.Marshal(order)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for error", func(t *testing.T) {
		mockDB.EXPECT().SellStock(&product).Return(nil, errors.New("error exist"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/product/consume", strings.NewReader(string(productJSON)))
		route.ServeHTTP(w, req)
		assert.Contains(t, w.Body.String(), "unable to consume stock")
		assert.Equal(t, w.Code, 500)

	})

	t.Run("Testing for successful consume", func(t *testing.T) {
		mockDB.EXPECT().SellStock(&product).Return(&order, nil).Times(1)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/product/consume", strings.NewReader(string(productJSON)))
		route.ServeHTTP(w, req)
		assert.Contains(t, w.Body.String(), string(orderJSON))
		assert.Equal(t, w.Code, http.StatusOK)

	})

}

func TestBulkUploadFromCsv(t *testing.T) {
	ctrl := gomock.NewController(t)
	//creates a new mock instance
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handlers.Handler{DB: mockDB}
	route := router.SetupRouter(h)

	file := [][]string{
		{"ma", "cbf87a9be799", "Foster-Harrell Table", "56"},
		{"dz", "e920c573f128", "Ramirez-Molina Granite Pizza", "47"},
	}
	fileJSON, err := json.Marshal(file)
	if err != nil {
		t.Fail()
	}
	t.Run("Testing for success upload", func(t *testing.T) {
		mockDB.EXPECT().BulkUpload(file).AnyTimes()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/product/bulkupdate", strings.NewReader(string(fileJSON)))
		route.ServeHTTP(w, req)
		assert.Equal(t, 200, 200)
	})
}
