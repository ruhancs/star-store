package dto

import "star_store/internal/domain/entity"

type InputCreateCartDto struct {
	ClientID  string             `json:"client_id"`
	CartItems []*entity.CartItem `json:"cart_items"`
}

type OutputCreateCartDto struct {
	ClientID  string             `json:"client_id"`
	CartItems []*entity.CartItem `json:"cart_items"`
	Total     float32            `json:"total"`
}

type InputInserItemOnCartDto struct {
	ClientID  string `json:"client_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	//Cart      *entity.Cart `json:"cart"`
}

type OutputInserItemOnCartDto struct {
	Items []*entity.CartItem `json:"items"`
	Total float32            `json:"total"`
}

type OutputGetCarByClientIDtDto struct {
	CartID    string             `json:"cart_id"`
	ClientID  string             `json:"client_id"`
	CartItems []*entity.CartItem `json:"cart_items"`
	Total     float32            `json:"total"`
}

type InputUpdateCartDto struct {
	Cart *entity.Cart `json:"cart"`
}

type OutputUpdateCartDto struct {
	Cart *entity.Cart `json:"cart"`
}

type OutputCheckoutUseCase struct {
	ClientName    string             `json:"client_name"`
	ClientZipCode string             `json:"client_zip_code"`
	Items         []*entity.CartItem `json:"items"`
	Total         float32            `json:"total"`
}
