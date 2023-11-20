package usecase_test

import (
	"errors"
	"star_store/internal/application/dto"
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"star_store/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInserItemCartUseCaseWithItemExistsOnCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cartRepository := mock_gateway.NewMockCartRepositoryInterface(ctrl)
	cartItemRepository := mock_gateway.NewMockCartItemRepository(ctrl)
	productRepository := mock_gateway.NewMockProductRepositoryInterface(ctrl)

	p1, _ := entity.NewProduct("p1", "1233", "JJ", "url", 20)
	cartItems := []*entity.CartItem{}
	cart := entity.NewCart("1234567", cartItems)
	cartItem1, _ := entity.NewCartItem(p1.Title, "1234567", cart.ID, 2, 20)

	cartRepository.EXPECT().Update(gomock.Any()).Return(nil)
	cartItemRepository.EXPECT().GetByCartID(cart.ID, p1.ID).Return(cartItem1,nil)
	
	inserItemCartUseCase := usecase.NewInsertItemOnCartUseCase(cartRepository, cartItemRepository, productRepository)
	input := dto.InputInserItemOnCartDto{
		ClientID: "1234567",
		ProductID: p1.ID,
		Quantity: cartItem1.Quantity,
		Cart: cart,
	}
	out,err := inserItemCartUseCase.Execute(input)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,1,len(out.Items))
	assert.Equal(t,40.0,out.Total)
}

func TestInserItemCartUseCaseWithItemNotExistsOnCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cartRepository := mock_gateway.NewMockCartRepositoryInterface(ctrl)
	cartItemRepository := mock_gateway.NewMockCartItemRepository(ctrl)
	productRepository := mock_gateway.NewMockProductRepositoryInterface(ctrl)

	p1, _ := entity.NewProduct("p1", "1233", "JJ", "url", 20)
	cartItems := []*entity.CartItem{}
	cart := entity.NewCart("1234567", cartItems)
	cartItem1, _ := entity.NewCartItem(p1.Title, "1234567", cart.ID, 2, 20)

	cartRepository.EXPECT().Update(gomock.Any()).Return(nil)
	cartItemRepository.EXPECT().GetByCartID(cart.ID, p1.ID).Return(nil,errors.New("not found"))
	cartItemRepository.EXPECT().Create(gomock.Any()).Return(nil)
	productRepository.EXPECT().Get(p1.ID).Return(p1,nil)
	

	inserItemCartUseCase := usecase.NewInsertItemOnCartUseCase(cartRepository, cartItemRepository, productRepository)
	input := dto.InputInserItemOnCartDto{
		ClientID: "1234567",
		ProductID: p1.ID,
		Quantity: cartItem1.Quantity,
		Cart: cart,
	}
	out,err := inserItemCartUseCase.Execute(input)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,1,len(out.Items))
	assert.Equal(t,40.0,out.Total)
}
