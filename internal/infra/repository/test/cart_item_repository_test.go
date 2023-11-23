package repository_test

import (
	"star_store/internal/domain/entity"
	"star_store/internal/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCartItem(t *testing.T) {
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	err := cartItemRepository.Create(cartItem1)
	
	assert.Nil(t,err)
	
	cartItemRepository.Delete(cartItem1.ID)
}

func TestListCartItems(t *testing.T) {
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItem2, _ := entity.NewCartItem(product2.Title, client.ID, cart.ID, 1, product2.Price)
	
	cartItemRepository.Create(cartItem1)
	cartItemRepository.Create(cartItem2)
	
	cartItems,err := cartItemRepository.List(cart.ID)
	
	assert.Nil(t,err)
	assert.NotNil(t,cartItems)
	assert.Equal(t,2,len(cartItems))

	cartItemRepository.Delete(cartItem1.ID)
	cartItemRepository.Delete(cartItem2.ID)
}

func TestGetCartItemByCartIDAndProductName(t *testing.T) {
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItemRepository.Create(cartItem1)

	cartItemFounded,err := cartItemRepository.GetByCartID(cartItem1.CartID,cartItem1.ProductName)

	assert.Nil(t,err)
	assert.NotNil(t,cartItemFounded)

	cartItemRepository.Delete(cartItem1.ID)
}

func TestNotFoundCartItemByCartIDAndProductName(t *testing.T) {
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItemRepository.Create(cartItem1)

	cartItemFounded,err := cartItemRepository.GetByCartID("notfound",cartItem1.ProductName)

	assert.NotNil(t,err)
	assert.Nil(t,cartItemFounded)

	cartItemRepository.Delete(cartItem1.ID)
}

func TestUpdateCartItem(t *testing.T) {
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItemRepository.Create(cartItem1)

	cartItem1.IncreaseQuantity(1)

	err := cartItemRepository.Update(cartItem1)
	assert.Nil(t,err)

	cartItemUpdated,err := cartItemRepository.GetByCartID(cartItem1.CartID,cartItem1.ProductName)
	assert.Nil(t,err)
	assert.Equal(t,2,cartItemUpdated.Quantity)
	assert.Equal(t,float32(40),cartItemUpdated.Total)
	cartItemRepository.Delete(cartItem1.ID)
}