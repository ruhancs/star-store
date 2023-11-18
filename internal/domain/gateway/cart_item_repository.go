package gateway

import "star_store/internal/domain/entity"

type CartItemRepository interface {
	Create(cartItem *entity.CartItem) error
	List(cartID string) ([]*entity.CartItem, error)
	Get(id string) (*entity.CartItem, error)
	Delete(id string) error
}
