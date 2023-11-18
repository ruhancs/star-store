package entity

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type CartItem struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ClientID     string  `json:"client_id"`
	CartID       string  `json:"cart_id"`
	Quantity     int     `json:"quantity"`
	Total        float64 `json:"total"`
}

func NewCartItem(productName, clientID, cartID string, quantity int, productPrice float64) (*CartItem, error) {
	cartItem := &CartItem{
		ID:           uuid.NewV4().String(),
		ProductName:  productName,
		ProductPrice: productPrice,
		ClientID:     clientID,
		Quantity:     quantity,
	}
	cartItem.Total = cartItem.CalculateTotal()
	err := cartItem.isValid()
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (c *CartItem) isValid() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *CartItem) CalculateTotal() float64 {
	total := c.ProductPrice * float64(c.Quantity)
	return total
}

func (c *CartItem) IncreaseQuantity(value int) {
	c.Quantity = c.Quantity + value
	c.Total = c.ProductPrice * float64(c.Quantity)
}

func (c *CartItem) DecreaseQuantity() {
	if c.Quantity > 0 {
		c.Quantity = c.Quantity - 1
		c.Total = c.ProductPrice * float64(c.Quantity)
	}
}
