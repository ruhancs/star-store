package entitytest

import (
	"star_store/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCart(t *testing.T) {
	ci1,_ := entity.NewCartItem("camisa","1236r7126","qwhg23", 3, 10)
	ci2,_ := entity.NewCartItem("camisa","1236r7126", "qwhg23", 3, 10)
	cartItems := []*entity.CartItem{ci1,ci2}
	cart := entity.NewCart("1236r7126",cartItems)

	assert.NotNil(t,cart)
	assert.Equal(t,"1236r7126",cart.ClientID)
	assert.Equal(t,cartItems[0],cart.CartItems[0])
	assert.Equal(t,len(cartItems),len(cart.CartItems))
	assert.Equal(t,float32(60),cart.Total)
}

func TestInsertItem(t *testing.T) {
	ci1,_ := entity.NewCartItem("camisa1","1236r7126","qwhg23", 1, 10)
	ci2,_ := entity.NewCartItem("camisa2","1236r7126", "qwhg23", 3, 10)
	ci3,_ := entity.NewCartItem("camisa3","1236r7126", "qwhg23", 1, 10)
	cartItems := []*entity.CartItem{ci1,ci2}
	cart := entity.NewCart("1236r7126",cartItems)

	assert.NotNil(t,cart)
	assert.Equal(t,cartItems[0].Quantity,cart.CartItems[0].Quantity)
	assert.Equal(t,1,cart.CartItems[0].Quantity)
	assert.Equal(t,float32(40),cart.Total)
	
	cart.InsertItem(ci1)
	assert.Equal(t,2,cart.CartItems[0].Quantity)
	assert.Equal(t,float32(50),cart.Total)
	
	cart.InsertItem(ci3)
	assert.Equal(t,len(cart.CartItems),3)
	assert.Equal(t,float32(60),cart.Total)
}
