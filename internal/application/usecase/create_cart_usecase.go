package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/entity"
	"star_store/internal/domain/gateway"
)

type CreateCartUseCase struct {
	CartRepository gateway.CartRepositoryInterface
}

func NewCreateCartUseCase(repo gateway.CartRepositoryInterface) *CreateCartUseCase{
	return &CreateCartUseCase{
		CartRepository: repo,
	}
}

func(u *CreateCartUseCase) Execute(input dto.InputCreateCartDto) (*dto.OutputCreateCartDto,error) {
	cart := entity.NewCart(input.ClientID,input.CartItems)
	err := u.CartRepository.Create(cart)
	if err != nil {
		return nil,err
	}
	output := &dto.OutputCreateCartDto{
		ClientID: cart.ClientID,
		CartItems: cart.CartItems,
		Total: cart.Total,
	}
	return output,nil
}