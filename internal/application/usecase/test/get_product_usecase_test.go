package usecase_test

import (
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"star_store/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductUseCase(t *testing.T) {
	p, _ := entity.NewProduct("p1", "qwe123", "Ju", "url.com", 10)
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	productMockRepo := mock_gateway.NewMockProductRepositoryInterface(ctrl)
	productMockRepo.EXPECT().Get(p.ID).Return(p, nil)
	getProductUseCase := usecase.NewGetProductUseCase(productMockRepo)

	out, err := getProductUseCase.Execute(p.ID)

	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.Equal(t, "p1", out.Item.Title)
	assert.Equal(t, "qwe123", out.Item.ZipCode)
	assert.Equal(t, "Ju", out.Item.Seller)
	assert.Equal(t, 10, out.Item.Price)
}
