package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/gateway"
)

type GetCartByClientIDUseCase struct {
	CartRepository gateway.CartRepositoryInterface
}

func NewGetCartByClientIDUseCase(repo gateway.CartRepositoryInterface) *GetCartByClientIDUseCase {
	return &GetCartByClientIDUseCase{
		CartRepository: repo,
	}
}

func (u *GetCartByClientIDUseCase) Execute(clientID string) (*dto.OutputGetCarByClientIDtDto, error) {
	cart, err := u.CartRepository.GetByUser(clientID)
	if err != nil {
		return nil, err
	}

	out := &dto.OutputGetCarByClientIDtDto{
		CartID:    cart.ID,
		ClientID:  cart.ClientID,
		CartItems: cart.CartItems,
		Total:     cart.Total,
	}
	return out, nil
}
