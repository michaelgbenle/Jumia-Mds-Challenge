package handlers

import (
	"github.com/golang/mock/gomock"

	mock_database "github.com/michaelgbenle/jumiaMds/database/mocks"
	"testing"
)

func TestGetProductBySku(t *testing.T) {
	ctrl := gomock.NewController(t)
	//creates a new mock instance
	mockDB := mock_database.NewMockDB(ctrl)
}
