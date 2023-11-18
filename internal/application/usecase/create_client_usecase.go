package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/entity"
	"star_store/internal/domain/gateway"
)

type CreateClietUseCase struct {
	ClientRepository gateway.ClientRepositoryInterface
}

func NewCreateClientUseCase(repo gateway.ClientRepositoryInterface) *CreateClietUseCase {
	return &CreateClietUseCase{
		ClientRepository: repo,
	}
}

func(u *CreateClietUseCase) Execute(input dto.InputCreateClientDto) (*dto.OutputCreateClientDto,error) {
	client,err := entity.NewClient(input.Name,input.ZipCode)
	if err != nil {
		return nil,err
	}

	err = u.ClientRepository.Create(client)

	output := &dto.OutputCreateClientDto{
		ID: client.ID,
		Name: client.Name,
		ZipCode: client.ZipCode,
	}

	return output,nil
}