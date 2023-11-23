package repository_test

import (
	"star_store/internal/domain/entity"
	"star_store/internal/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {
	cartRepository := repository.NewCartRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItem2, _ := entity.NewCartItem(product2.Title, client.ID, cart.ID, 1, product2.Price)
	cart.InsertItem(cartItem1)
	cart.InsertItem(cartItem2)
	err := cartRepository.Create(cart)

	assert.Nil(t, err)
	cartRepository.Delete(cart.ID)
}

func TestUpdateCart(t *testing.T) {
	cartRepository := repository.NewCartRepository(session)
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItem2, _ := entity.NewCartItem(product2.Title, client.ID, cart.ID, 1, product2.Price)
	cart.InsertItem(cartItem1)
	cartItemRepository.Create(cartItem1)
	err := cartRepository.Create(cart)
	assert.Nil(t, err)

	cart.InsertItem(cartItem2)
	cartItemRepository.Create(cartItem2)
	err = cartRepository.Update(cart)
	assert.Nil(t, err)

	cartFounded, err := cartRepository.GetByID(cart.ID)
	assert.Nil(t, err)
	assert.NotNil(t, cartFounded)
	assert.Equal(t, cart.ID, cartFounded.ID)
	assert.Equal(t, cart.ClientID, cartFounded.ClientID)
	assert.Equal(t, 2, len(cartFounded.CartItems))
	assert.Equal(t, float32(30), cartFounded.Total)

	cartRepository.Delete(cart.ID)
	cartItemRepository.Delete(cartItem1.ID)
	cartItemRepository.Delete(cartItem2.ID)
}

func TestGetCartByID(t *testing.T) {
	cartRepository := repository.NewCartRepository(session)
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItem2, _ := entity.NewCartItem(product2.Title, client.ID, cart.ID, 1, product2.Price)

	cartItemRepository.Create(cartItem1)
	cartItemRepository.Create(cartItem2)
	cart.InsertItem(cartItem1)
	cart.InsertItem(cartItem1)
	cart.InsertItem(cartItem2)
	cartRepository.Create(cart)

	cartFounded, err := cartRepository.GetByID(cart.ID)

	assert.Nil(t, err)
	assert.NotNil(t, cartFounded)
	assert.Equal(t, cart.ID, cartFounded.ID)
	assert.Equal(t, cart.ClientID, cartFounded.ClientID)
	assert.Equal(t, 2, len(cartFounded.CartItems))
	assert.Equal(t, float32(50), cartFounded.Total)

	cartRepository.Delete(cart.ID)
	cartItemRepository.Delete(cartItem1.ID)
	cartItemRepository.Delete(cartItem2.ID)
}

func TestGetCartByClientID(t *testing.T) {
	cartRepository := repository.NewCartRepository(session)
	cartItemRepository := repository.NewCartItemRepository(session)
	cart := entity.NewCart(client.ID, []*entity.CartItem{})
	cartItem1, _ := entity.NewCartItem(product.Title, client.ID, cart.ID, 1, product.Price)
	cartItem2, _ := entity.NewCartItem(product2.Title, client.ID, cart.ID, 1, product2.Price)

	cartItemRepository.Create(cartItem1)
	cartItemRepository.Create(cartItem2)
	cart.InsertItem(cartItem1)
	cart.InsertItem(cartItem1)
	cart.InsertItem(cartItem2)
	cartRepository.Create(cart)

	cartFounded, err := cartRepository.GetByUser(cart.ClientID)

	assert.Nil(t, err)
	assert.NotNil(t, cartFounded)
	assert.Equal(t, cart.ID, cartFounded.ID)
	assert.Equal(t, cart.ClientID, cartFounded.ClientID)
	assert.Equal(t, 2, len(cartFounded.CartItems))
	assert.Equal(t, float32(50), cartFounded.Total)

	cartRepository.Delete(cart.ID)
	cartItemRepository.Delete(cartItem1.ID)
	cartItemRepository.Delete(cartItem2.ID)

	cartItemRepository.Delete("ed70c848-b499-487a-8efe-6f1cf6371e92")
}
