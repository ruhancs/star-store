package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	PurchaseID string  `json:"purchase_id" valid:"required"`
	ClientID   string  `json:"client_id" valid:"required"`
	ClientName string  `json:"client_name" valid:"required"`
	TotalToPay float32 `json:"total_to_pay" valid:"required"`
	CreditCard string  `json:"credit_card" valid:"required,stringlength(16|16)"`
	Date       string  `json:"date"`
}

func NewTransaction(clientId, clientName, creditCard string, totalToPay float32) (*Transaction, error) {
	transaction := &Transaction{
		PurchaseID: uuid.NewV4().String(),
		ClientID:   clientId,
		ClientName: clientName,
		TotalToPay: totalToPay,
		CreditCard: creditCard,
		Date:       time.Now().UTC().String(),
	}

	err := transaction.isValid()
	if err != nil {
		if err.Error() == "credit_card: 1235 does not validate as stringlength(16|16)" {
			return nil, errors.New("invalid credit card number")
		}
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	return nil
}
