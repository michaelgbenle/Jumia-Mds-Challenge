package test

import (
	"github.com/golang/mock/gomock"
	mockdatabase "github.com/michaelgbenle/jumiaMds/database/mocks"
	"github.com/michaelgbenle/jumiaMds/models"
	"github.com/michaelgbenle/jumiaMds/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProductBySku(t *testing.T) {
	ctrl := gomock.NewController(t)
	//creates a new mock instance
	mockDB := mockdatabase.NewMockDB(ctrl)
	//h := handlers.HandleConstruct()
	route := router.SetupRouter()
	product := models.Product{
		Name:    "Samsung Phone",
		Sku:     "UYUT-879847564793-PO",
		Stock:   2,
		Country: "ke",
	}
	mockDB.EXPECT().GetProductSku(product.Sku, product.Country).Return(product, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "api/v1/product", strings.NewReader(""))
	route.ServeHTTP(w, req)
	assert.Contains(t, w.Body.String(), product)
	assert.Equal(t, w.Code, http.StatusOK)

}
