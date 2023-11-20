package usecase_test

import (
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"star_store/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetCartByClientIDUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cartRepository := mock_gateway.NewMockCartRepositoryInterface(ctrl)
	client,_ := entity.NewClient("C1","12345567")
	cartItem1,_ := entity.NewCartItem("P1","124345","1232435",1,20)
	cartItem2,_ := entity.NewCartItem("P2","124345","1232435",2,10)
	cartItems := []*entity.CartItem{cartItem1,cartItem2}
	cart := entity.NewCart(client.ID,cartItems)
	cartRepository.EXPECT().GetByUser(client.ID).Return(cart,nil)

	getCartByClientIDUseCase := usecase.NewGetCartByClientIDUseCase(cartRepository)
	out,err := getCartByClientIDUseCase.Execute(client.ID)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,cart.ClientID,out.ClientID)
	assert.Equal(t,cart.Total,out.Total)
	assert.Equal(t,cart.CartItems,out.CartItems)
}