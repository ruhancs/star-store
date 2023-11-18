package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/gateway"
)

type GetClientUseCase struct {
	ClientRepository gateway.ClientRepositoryInterface
}

func NewGetClientUseCase(repo gateway.ClientRepositoryInterface) *GetClientUseCase {
	return &GetClientUseCase{
		ClientRepository: repo,
	}
}

func(u *GetClientUseCase) Execute(id string) (*dto.OutputGetClientDto,error) {
	client,err := u.ClientRepository.Get(id)
	if err != nil {
		return nil,err
	}

	output := dto.OutputGetClientDto{
		Client: &client,
	}

	return &output,nil
} 