package repository_test

import (
	"star_store/internal/domain/entity"
	"star_store/internal/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var transactionRepository = repository.NewTransactionRepository(session)
var transaction, _ = entity.NewTransaction(client.ID, client.Name, "1234123412341234", 20)
var transaction2, _ = entity.NewTransaction(client.ID, client.Name, "1234123412341234", 10)

func TestCreateTransaction(t *testing.T) {
	err := transactionRepository.Create(transaction)

	assert.Nil(t, err)

	transactionRepository.Delete(transaction.PurchaseID)
}

func TestGetClientTransactions(t *testing.T) {
	transactionRepository.Create(transaction)
	transactionRepository.Create(transaction2)

	transactions, err := transactionRepository.GetByClientID(client.ID)

	assert.Nil(t, err)
	assert.NotNil(t, transactions)
	assert.Equal(t, 2, len(transactions))
	assert.Equal(t, client.Name, transactions[0].ClientName)
	assert.Equal(t, float32(10), transactions[0].TotalToPay)
	assert.Equal(t, float32(20), transactions[1].TotalToPay)

	transactionRepository.Delete(transaction.PurchaseID)
	transactionRepository.Delete(transaction2.PurchaseID)
}
