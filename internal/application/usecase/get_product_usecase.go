package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/gateway"
)

type GetProductUseCase struct {
	ProductRepository gateway.ProductRepositoryInterface
}

func NewGetProductUseCase(repo gateway.ProductRepositoryInterface) *GetProductUseCase{
	return &GetProductUseCase{
		ProductRepository: repo,
	}
}

func(u *GetProductUseCase) Execute(id string) (*dto.OutputGetProductDto,error) {
	product,err := u.ProductRepository.Get(id)
	if err != nil {
		return nil,err
	}

	output := &dto.OutputGetProductDto{
		Item: product,
	}

	return output,nil
}