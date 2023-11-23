package usecase

import (
	"errors"
	"star_store/internal/application/dto"
	"star_store/internal/domain/entity"
	"star_store/internal/domain/gateway"
)

type CheckoutUseCase struct {
	CartRepository gateway.CartRepositoryInterface
	ClientRepository gateway.ClientRepositoryInterface
}

func NewCheckoutUseCase(cartRepo gateway.CartRepositoryInterface, clientRepo gateway.ClientRepositoryInterface) *CheckoutUseCase {
	return &CheckoutUseCase{
		CartRepository: cartRepo,
		ClientRepository: clientRepo,
	}
} 

func(u *CheckoutUseCase) Execute(cart *entity.Cart) (*dto.OutputCheckoutUseCase,error) {
	err := u.CartRepository.Create(cart)
	if err != nil {
		return nil,err
	}
	client,err := u.ClientRepository.Get(cart.ClientID)
	if err != nil {
		return nil,errors.New("client not found")
	}

	output := &dto.OutputCheckoutUseCase{
		Items: cart.CartItems,
		ClientName: client.Name,
		ClientZipCode: client.ZipCode,
		Total: cart.Total,
	}

	return output,nil
}