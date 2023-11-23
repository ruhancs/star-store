package gateway

import "star_store/internal/domain/entity"

type TransactionRepositoryInterface interface {
	Create(transaction *entity.Transaction) error
	GetByClientID(clientID string) ([]*entity.Transaction,error)
}