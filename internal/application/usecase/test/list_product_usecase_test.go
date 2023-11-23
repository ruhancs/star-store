package usecase_test

import (
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"star_store/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListProductUseCase(t *testing.T) {
	p1,_ := entity.NewProduct("p1","qwe123","Ju","url.com",10)
	p2,_ := entity.NewProduct("p2","jash7","Ju","url.com",20)
	products := []*entity.Product{p1,p2}
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	productMockRepo := mock_gateway.NewMockProductRepositoryInterface(ctrl)
	productMockRepo.EXPECT().List().Return(products,nil)
	listProductUseCase := usecase.NewListProductUseCase(productMockRepo)

	out,err := listProductUseCase.Execute()

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,"p1",out.Items[0].Title)
	assert.Equal(t,"qwe123",out.Items[0].ZipCode)
	assert.Equal(t,"Ju",out.Items[0].Seller)
	assert.Equal(t,float32(10),out.Items[0].Price)
	assert.Equal(t,"p2",out.Items[1].Title)
	assert.Equal(t,"jash7",out.Items[1].ZipCode)
	assert.Equal(t,"Ju",out.Items[1].Seller)
	assert.Equal(t,float32(20),out.Items[1].Price)
}