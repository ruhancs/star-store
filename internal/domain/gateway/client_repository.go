package gateway

import "star_store/internal/domain/entity"

type ClientRepositoryInterface interface {
	Create(client *entity.Client) error
	Get(id string) (*entity.Client, error)
}
