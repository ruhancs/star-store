package repository

import (
	"star_store/internal/domain/entity"

	"github.com/gocql/gocql"
)

type CartRepository struct {
	DB *gocql.Session
}

func NewCartRepository(db *gocql.Session) *CartRepository{
	return &CartRepository{
		DB: db,
	}
}

func(c *CartRepository) Create(cart *entity.Cart) error {
	cartItemsIDs := cart.GetCartItemsIDs()
	err := c.DB.Query(`
	INSERT INTO carts(id, client_id, total, cart_items) VALUES (?, ?, ?, ?);`,
	cart.ID,cart.ClientID,cart.Total,cartItemsIDs).Exec()
	if err != nil {
		return err
	}
	
	return nil
}

func(c *CartRepository) Update(cart *entity.Cart) error {
	cartItemsIDs := cart.GetCartItemsIDs()
	err := c.DB.Query(`
		UPDATE carts SET total=?, cart_items=? WHERE id=?;`,
			cart.Total, cartItemsIDs, cart.ID).Exec()
	if err != nil {
		return err
	}

	return nil
}

func(p *CartRepository) GetByID(id string) (*entity.Cart,error) {
	var cart entity.Cart
	var cartItemsIDs []string
	err := p.DB.Query("SELECT * from carts WHERE id=?;",id).Scan(&cart.ID,&cartItemsIDs,&cart.ClientID,&cart.Total)
	if err != nil {
		return nil,err
	}
	for _,itemID := range cartItemsIDs {
		var cartItem entity.CartItem
		err := p.DB.Query(`
			SELECT * from cart_items WHERE id=?;`,itemID).Scan(
				&cartItem.ID,&cartItem.CartID,&cartItem.ClientID,&cartItem.ProductName,&cartItem.ProductPrice,&cartItem.Quantity,&cartItem.Total)
		if err != nil {
			return nil,err
		}
		cart.CartItems = append(cart.CartItems, &cartItem)
	}

	return &cart,nil
}

func(p *CartRepository) GetByUser(userID string) (*entity.Cart,error) {
	var cart entity.Cart
	var cartItemsIDs []string
	err := p.DB.Query("SELECT * from carts WHERE client_id=? ALLOW FILTERING;",userID).Scan(&cart.ID,&cartItemsIDs,&cart.ClientID,&cart.Total)
	if err != nil {
		return nil,err
	}
	for _,itemID := range cartItemsIDs {
		var cartItem entity.CartItem
		err := p.DB.Query(`
			SELECT * from cart_items WHERE id=?;`,itemID).Scan(
				&cartItem.ID,&cartItem.CartID,&cartItem.ClientID,&cartItem.ProductName,&cartItem.ProductPrice,&cartItem.Quantity,&cartItem.Total)
		if err != nil {
			return nil,err
		}
		cart.CartItems = append(cart.CartItems, &cartItem)
	}

	return &cart,nil
}

func(c *CartRepository) Delete(id string) error {
	err := c.DB.Query(`Delete FROM carts WHERE id=?`,id).Exec()
	if err != nil {
		return err
	}

	return nil
}