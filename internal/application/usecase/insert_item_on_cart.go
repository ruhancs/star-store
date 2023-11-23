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
		ProductRepository:  productRepo,
	}
}

func (u *InsertItemOnCartUseCase) Execute(input dto.InputInserItemOnCartDto, cart *entity.Cart, cartItem *entity.CartItem) (*dto.OutputInserItemOnCartDto, error) {
	_, err := u.CartRepository.GetByID(cart.ID)
	if err != nil {
		err := u.CartRepository.Create(cart)
		if err != nil {
			return nil, err
		}
		err = u.CartItemRepository.Create(cartItem)
		if err != nil {
			//user unit of work
			err = u.CartItemRepository.Create(cartItem)
		}
		return &dto.OutputInserItemOnCartDto{
			Items: cart.CartItems,
			Total: cart.Total,
		}, nil
	}

	cartItemExist, err := u.CartItemRepository.GetByCartID(cart.ID, cartItem.ProductName)
	if err != nil {
		err = u.CartItemRepository.Create(cartItem)
		if err != nil {
			//user unit of work
			err = u.CartItemRepository.Create(cartItem)
		}
		err := u.CartRepository.Update(cart)
		if err != nil {
			return nil, err
		}
		return &dto.OutputInserItemOnCartDto{
			Items: cart.CartItems,
			Total: cart.Total,
		}, nil
	}

	cartItemUpdated := &entity.CartItem{
		ID: cartItemExist.ID,
		ProductName: cartItemExist.ProductName,
		ProductPrice: cartItemExist.ProductPrice,
		ClientID: cartItemExist.ClientID,
		CartID: cartItemExist.CartID,
		Quantity: cartItemExist.Quantity + cartItem.Quantity,
		Total: cartItemExist.CalculateTotal() + cartItem.Total,
	}
	u.CartItemRepository.Update(cartItemUpdated)
	u.CartRepository.Update(cart)
	return &dto.OutputInserItemOnCartDto{
		Items: cart.CartItems,
		Total: cart.Total,
	}, nil
}
