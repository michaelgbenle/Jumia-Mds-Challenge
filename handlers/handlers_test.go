package handlers

import (
	"github.com/golang/mock/gomock"
	"github.com/michaelgbenle/jumiaMds/models"
	"github.com/michaelgbenle/jumiaMds/router"

	mock_database "github.com/michaelgbenle/jumiaMds/database/mocks"
	"testing"
)

func TestGetProductBySku(t *testing.T) {
	ctrl := gomock.NewController(t)
	//creates a new mock instance
	mockDB := mock_database.NewMockDB(ctrl)
	h := &handler{
		DB: mockDB,
	}
	route := router.SetupRouter()
	product := models.Product{
		Name:    "Samsung Phone",
		Sku:     "UYUT-879847564793-PO",
		Stock:   2,
		Country: "ke",
	}
	mockDB.GetProductSku(product.Sku, product.Country)
}
