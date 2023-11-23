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
	cart.InsertItem(cartItem1)

	cartRepository.EXPECT().GetByID(cart.ID).Return(cart,nil)
	cartRepository.EXPECT().Update(gomock.Any()).Return(nil)
	cartItemRepository.EXPECT().GetByCartID(cart.ID, p1.Title).Return(cartItem1,nil)
	cartItemRepository.EXPECT().Update(gomock.Any()).Return(nil)
	
	inserItemCartUseCase := usecase.NewInsertItemOnCartUseCase(cartRepository, cartItemRepository, productRepository)
	input := dto.InputInserItemOnCartDto{
		ClientID: "1234567",
		ProductID: p1.ID,
		Quantity: cartItem1.Quantity,
	}
	out,err := inserItemCartUseCase.Execute(input,cart,cartItem1)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,1,len(out.Items))
	assert.Equal(t,float32(40),out.Total)
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
	cart.InsertItem(cartItem1)

	cartRepository.EXPECT().GetByID(cart.ID).Return(cart,nil)
	cartRepository.EXPECT().Update(gomock.Any()).Return(nil)
	cartItemRepository.EXPECT().GetByCartID(cart.ID, p1.Title).Return(nil,errors.New("not found"))
	cartItemRepository.EXPECT().Create(gomock.Any()).Return(nil)	

	inserItemCartUseCase := usecase.NewInsertItemOnCartUseCase(cartRepository, cartItemRepository, productRepository)
	input := dto.InputInserItemOnCartDto{
		ClientID: "1234567",
		ProductID: p1.ID,
		Quantity: cartItem1.Quantity,
	}
	out,err := inserItemCartUseCase.Execute(input,cart,cartItem1)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,1,len(out.Items))
	assert.Equal(t,float32(40),out.Total)
}
