package gateway

import "star_store/internal/domain/entity"


type ProductRepositoryInterface interface {
	Create(product *entity.Product) error
	List() ([]*entity.Product,error)
	Get(id string) (*entity.Product,error)
}