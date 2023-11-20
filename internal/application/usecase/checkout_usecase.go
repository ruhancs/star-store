package usecase

import (
	"errors"
	"star_store/internal/application/dto"
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

func(u *CheckoutUseCase) Execute(cartID string) (*dto.OutputCheckoutUseCase,error) {
	cart,err := u.CartRepository.GetByID(cartID)
	if err != nil {
		return nil,errors.New("cart not found")
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