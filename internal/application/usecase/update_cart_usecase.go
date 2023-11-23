package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/gateway"
)

type UpdateCartUseCase struct {
	CartRepository gateway.CartRepositoryInterface
}

func NewUpdateCartUseCase(repo gateway.CartRepositoryInterface) *UpdateCartUseCase{
	return &UpdateCartUseCase{
		CartRepository: repo,
	}
}

func(u *UpdateCartUseCase) Execute(input dto.InputUpdateCartDto) (*dto.OutputUpdateCartDto,error) {
	err := u.CartRepository.Update(input.Cart)
	if err != nil {
		return nil,err
	}

	output := &dto.OutputUpdateCartDto{
		Cart: input.Cart,
	}
	return output,nil
}