package gateway

import "star_store/internal/domain/entity"

type CartItemRepository interface {
	Create(cartItem *entity.CartItem) error
	List(cartID string) ([]*entity.CartItem, error)
	//Get(id string) (*entity.CartItem, error)
	GetByCartID(cartId, productName string) (*entity.CartItem, error)
	Update(cartItem *entity.CartItem) error
	Delete(id string) error
}
