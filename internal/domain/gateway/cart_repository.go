package gateway

import "star_store/internal/domain/entity"

type CartRepositoryInterface interface {
	Create(cart *entity.Cart) error
	GetByID(id string) (*entity.Cart,error)
	GetByUser(userID string) (*entity.Cart,error)
	Update(cart *entity.Cart) error
}