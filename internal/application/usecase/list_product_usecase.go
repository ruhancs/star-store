package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/gateway"
)

type ListProductUseCase struct {
	ProductRepository gateway.ProductRepositoryInterface
}

func NewListProductUseCase(repo gateway.ProductRepositoryInterface) *ListProductUseCase {
	return &ListProductUseCase{
		ProductRepository: repo,
	}
}

func(u *ListProductUseCase) Execute() (*dto.OutputListProductDto,error){
	products,err := u.ProductRepository.List()
	if err != nil {
		return nil,err
	}
	output := &dto.OutputListProductDto{
		Items: products,
	}
	
	return output,nil
}