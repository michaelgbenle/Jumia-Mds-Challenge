package test

import (
	"encoding/json"
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

	product2 := models.Product{
		Name:    "laptop",
		Sku:     "e920c573f128",
		Stock:   2,
		Country: "gh",
	}
	product2JSON, err := json.Marshal(product2)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for successful request", func(t *testing.T) {
		mockDB.EXPECT().GetProductSku("cbf87a9be799", "ma").Return(&product, nil).Times(1)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/product?sku=cbf87a9be799&country=ma", strings.NewReader(string(productJSON)))
		route.ServeHTTP(w, req)
		assert.Contains(t, w.Body.String(), "ma")
		assert.Equal(t, w.Code, http.StatusOK)

	})

	t.Run("Testing for error", func(t *testing.T) {
		mockDB.EXPECT().GetProductSku("e920c573f128", "gh").Return(nil, err)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/product?", strings.NewReader(string(product2JSON)))
		route.ServeHTTP(w, req)
		//assert.Contains(t, w.Body.String(), "empty")
		assert.Equal(t, w.Code, 400)

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
		Quantity:  3,
	}

	t.Run("Testing for successful consume", func(t *testing.T) {
		mockDB.EXPECT().SellStock(product).Return(&order, nil).Times(1)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/product/consume", strings.NewReader(string(productJSON)))
		route.ServeHTTP(w, req)
		assert.Contains(t, w.Body.String(), order)
		assert.Equal(t, w.Code, http.StatusOK)

	})

}
