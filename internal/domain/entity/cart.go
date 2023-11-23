package entity

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	ID        string      `json:"id"`
	ClientID  string      `json:"client_id"`
	CartItems []*CartItem `json:"cart_items"`
	Total     float32     `json:"total"`
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

func (c *Cart) CalculateTotal() float32 {
	var total = float32(0.0)
	for _, item := range c.CartItems {
		total += item.Total 
	}
	return total
}

func(c *Cart) InsertItem(item *CartItem) {
	for index,i := range c.CartItems {
		if i.ProductName == item.ProductName {
			fmt.Println(c.CartItems[index].Quantity)
			i.Quantity = item.Quantity + i.Quantity
			c.CartItems[index].Quantity = i.Quantity
			c.CartItems[index].Total = c.CartItems[index].CalculateTotal()
			c.Total = c.CalculateTotal()
			fmt.Println(c.CartItems[index].Quantity)
			return
		}
	}
	c.CartItems = append(c.CartItems, item)
	c.Total = c.CalculateTotal()
	return
}

func(c *Cart) GetCartItemsIDs() []string {
	var ids []string
	for _,item := range c.CartItems {
		ids = append(ids, item.ID)
	}
	return ids
}
