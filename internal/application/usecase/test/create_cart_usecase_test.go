package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"star_store/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCreatCartUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cartRepository := mock_gateway.NewMockCartRepositoryInterface(ctrl)
	cartItem1,_ := entity.NewCartItem("P1","124345","1232435",1,20)
	cartItem2,_ := entity.NewCartItem("P2","124345","1232435",2,10)
	cartItems := []*entity.CartItem{cartItem1,cartItem2}
	cart := entity.NewCart("1234567",cartItems)
	cartRepository.EXPECT().Create(gomock.Any()).Return(nil)

	createCartUseCase := usecase.NewCreateCartUseCase(cartRepository)
	input := dto.InputCreateCartDto{
		ClientID: cart.ClientID,
		CartItems: cartItems,
	}
	out,err := createCartUseCase.Execute(input)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,cart.ClientID,out.ClientID)
	assert.Equal(t,cart.Total,out.Total)
	assert.Equal(t,float32(40),out.Total)
	assert.Equal(t,cart.CartItems[0],out.CartItems[0])
}