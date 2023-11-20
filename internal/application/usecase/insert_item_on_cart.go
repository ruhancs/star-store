package usecase

import (
	"star_store/internal/application/dto"
	"star_store/internal/domain/entity"
	"star_store/internal/domain/gateway"
)

type InsertItemOnCartUseCase struct {
	CartRepository     gateway.CartRepositoryInterface
	CartItemRepository gateway.CartItemRepository
	ProductRepository  gateway.ProductRepositoryInterface
}

func NewInsertItemOnCartUseCase(
	cartRepo gateway.CartRepositoryInterface,
	cartItemRepo gateway.CartItemRepository,
	productRepo gateway.ProductRepositoryInterface,
) *InsertItemOnCartUseCase {
	return &InsertItemOnCartUseCase{
		CartRepository:     cartRepo,
		CartItemRepository: cartItemRepo,
		ProductRepository: productRepo,
	}
}

func (u *InsertItemOnCartUseCase) Execute(input dto.InputInserItemOnCartDto) (*dto.OutputInserItemOnCartDto,error) {
	cartItem, err := u.CartItemRepository.GetByCartID(input.Cart.ID, input.ProductID)
	if err != nil {
		product,err := u.ProductRepository.Get(input.ProductID)
		if err != nil {
			return nil, err
		}
		cartItem, err := entity.NewCartItem(product.Title,input.ClientID,input.Cart.ID,input.Quantity,float64(product.Price))
		if err != nil {
			return nil, err
		}
		err = u.CartItemRepository.Create(cartItem)
		if err != nil {
			return nil, err
		}
		input.Cart.InsertItem(cartItem)
		u.CartRepository.Update(input.Cart)
		return &dto.OutputInserItemOnCartDto{
			Items: input.Cart.CartItems,
			Total: input.Cart.Total,
			},nil
		}
		
	input.Cart.InsertItem(cartItem)
	u.CartRepository.Update(input.Cart)
	return &dto.OutputInserItemOnCartDto{
		Items: input.Cart.CartItems,
		Total: input.Cart.Total,
	},nil
}
