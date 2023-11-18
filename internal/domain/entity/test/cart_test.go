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
	assert.Equal(t,60.0,cart.Total)
}
