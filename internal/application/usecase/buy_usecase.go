package usecase

import (
	"errors"
	"star_store/internal/application/dto"
	"star_store/internal/domain/entity"
	"star_store/internal/domain/gateway"
)

type BuyUseCase struct {
	TransactionRepository gateway.TransactionRepositoryInterface
	ClientRepository      gateway.ClientRepositoryInterface
}

func NewBuyUseCase(
	transactionRepo gateway.TransactionRepositoryInterface,
	clientRepo gateway.ClientRepositoryInterface,
) *BuyUseCase {
	return &BuyUseCase{
		TransactionRepository: transactionRepo,
		ClientRepository:      clientRepo,
	}
}

func (u *BuyUseCase) Execute(input dto.InputBuyUseCaseDto) (*dto.OutputBuyUseCaseDto, error) {
	client, err := u.ClientRepository.Get(input.ClientID)
	if err != nil {
		return nil, errors.New("client not found")
	}
	transaction, err := entity.NewTransaction(client.ID, client.Name, input.CardNumber, input.Value)
	if err != nil {
		return nil, err
	}

	err = u.TransactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}

	output := &dto.OutputBuyUseCaseDto{
		ClientID:         client.ID,
		ClientName:       client.Name,
		TotalToPay:       input.Value,
		CreditCardNumber: input.CardNumber,
	}

	return output, nil
}
