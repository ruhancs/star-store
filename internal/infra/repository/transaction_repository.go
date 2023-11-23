package repository

import (
	"star_store/internal/domain/entity"

	"github.com/gocql/gocql"
)

type TransactionRepository struct {
	DB *gocql.Session
}

func NewTransactionRepository(db *gocql.Session) *TransactionRepository{
	return &TransactionRepository{
		DB: db,
	}
}

func(c *TransactionRepository)Create(transaction *entity.Transaction) error {
	err := c.DB.Query(`
		INSERT INTO transactions(purchase_id, client_id, client_name, total_to_pay, credit_card, date) VALUES (?, ?, ?, ?,?,?);`,
			transaction.PurchaseID,transaction.ClientID,transaction.ClientName,transaction.TotalToPay,transaction.CreditCard,transaction.Date).Exec()
	if err != nil {
		return err
	}

	return nil
}

func(p *TransactionRepository) GetByClientID(clientID string) ([]*entity.Transaction,error) {
	var transactions []*entity.Transaction
	rows := p.DB.Query("SELECT * from transactions WHERE client_id=? ALLOW FILTERING;",clientID).Iter()
	defer rows.Close()
	scanner := rows.Scanner()
	for scanner.Next() {
		var transaction entity.Transaction
		err := scanner.Scan(&transaction.PurchaseID,&transaction.ClientID,&transaction.ClientName,&transaction.CreditCard,&transaction.Date,&transaction.TotalToPay)
		if err != nil {
			return nil,err
		}
		transactions = append(transactions, &transaction)
	}

	return transactions,nil
}

func(c *TransactionRepository) Delete(purchaseID string) error {
	err := c.DB.Query(`Delete FROM transactions WHERE purchase_id=?`,purchaseID).Exec()
	if err != nil {
		return err
	}

	return nil
}