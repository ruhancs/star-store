package usecase

import (
	"errors"
	"star_store/internal/application/dto"
	"star_store/internal/domain/entity"
	"star_store/internal/domain/gateway"
)

type CreateProducteUseCase struct {
	ProductRepository gateway.ProductRepositoryInterface
}

func NewCreateProductUseCase(repo gateway.ProductRepositoryInterface) *CreateProducteUseCase{
	return &CreateProducteUseCase{
		ProductRepository: repo,
	}
}

func(u *CreateProducteUseCase) Execute(input dto.InputCreateProductDto) (*dto.OutputCreateProductDto,error) {
	product,err := entity.NewProduct(input.Title,input.ZipCode,input.Seller,input.ThumbnailHD,input.Price)
	if err != nil {
		return nil,errors.New("invalid product data: " +err.Error())
	}

	err = u.ProductRepository.Create(product)
	if err != nil {
		return nil,err
	}

	output := &dto.OutputCreateProductDto{
		ID: product.ID,
		Title: product.Title,
		Price: product.Price,
		ZipCode: product.ZipCode,
		Seller: product.Seller,
		ThumbnailHD: product.ThumbnailHD,
	}

	return output,nil
} 