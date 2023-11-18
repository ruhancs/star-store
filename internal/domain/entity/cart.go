package entity

import uuid "github.com/satori/go.uuid"

type Cart struct {
	ID        string      `json:"id"`
	ClientID  string      `json:"client_id"`
	CartItems []*CartItem `json:"cart_items"`
	Total     float64     `json:"total"`
}

func NewCart(clientID string, cartItems []*CartItem) *Cart {
	cart := &Cart{
		ID:        uuid.NewV4().String(),
		ClientID:  clientID,
		CartItems: cartItems,
	}
	total := cart.CalculateTotal()
	cart.Total = total
	return cart
}

func (c *Cart) CalculateTotal() float64 {
	var total = 0.0
	for _, item := range c.CartItems {
		total += item.Total
	}
	return total
}
