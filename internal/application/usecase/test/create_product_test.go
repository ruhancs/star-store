package usecase_test

import (
	"star_store/internal/application/dto"
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func TestNewCreateProducteUseCase(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	productMockRepo := mock_gateway.NewMockProductRepositoryInterface(ctrl)
	productMockRepo.EXPECT().Create(gomock.Any()).AnyTimes()
	createProductUseCase := usecase.NewCreateProducteUseCase(productMockRepo)

	input := dto.InputCreateProductDto{
		Title: "P1",
		Price: 10,
		ZipCode: "produc_1",
		Seller: "Joao",
		ThumbnailHD: "url.com",
	}
	output,err := createProductUseCase.Execute(input)

	assert.Nil(t,err)
	assert.NotNil(t,output)
	assert.Equal(t,"P1",output.Title)
	assert.Equal(t,10,output.Price)
}