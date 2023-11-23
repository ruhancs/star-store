package repository

import (
	"star_store/internal/domain/entity"

	"github.com/gocql/gocql"
)

type CartItemRepository struct {
	DB *gocql.Session
}

func NewCartItemRepository(db *gocql.Session) *CartItemRepository{
	return &CartItemRepository{
		DB: db,
	}
}

func(c *CartItemRepository) Create(cartItem *entity.CartItem) error {
	err := c.DB.Query(`
		INSERT INTO cart_items(id, cart_id, client_id, product_name, product_price, quantity, total) VALUES (?, ?, ?, ?,?,?,?);`,
			cartItem.ID,cartItem.CartID,cartItem.ClientID,cartItem.ProductName,cartItem.ProductPrice,cartItem.Quantity,cartItem.Total).Exec()
	if err != nil {
		return err
	}

	return nil
}

func(p *CartItemRepository) List(cartID string) ([]*entity.CartItem, error) {
	var cartItems []*entity.CartItem
	rows := p.DB.Query("SELECT * from cart_items WHERE cart_id=? ALLOW FILTERING;",cartID).Iter()
	defer rows.Close()
	scanner := rows.Scanner()
	for scanner.Next() {
		var cartItem entity.CartItem
		err := scanner.Scan(&cartItem.ID,&cartItem.CartID,&cartItem.ClientID,&cartItem.ProductName,&cartItem.ProductPrice,&cartItem.Quantity,&cartItem.Total)
		if err != nil {
			return nil,err
		}
		cartItems = append(cartItems, &cartItem)
	}

	return cartItems,nil
}

func(p *CartItemRepository) GetByCartID(cartId, productName string) (*entity.CartItem, error) {
	var cartItem entity.CartItem
	err := p.DB.Query(`
		SELECT * from cart_items WHERE cart_id=? AND product_name=? ALLOW FILTERING;`,cartId,productName).Scan(
			&cartItem.ID,&cartItem.CartID,&cartItem.ClientID,&cartItem.ProductName,&cartItem.ProductPrice,&cartItem.Quantity,&cartItem.Total)
	if err != nil {
		return nil,err
	}

	return &cartItem,nil
}

func(c *CartItemRepository) Update(cartItem *entity.CartItem) error {
	err := c.DB.Query(`
		UPDATE cart_items SET total=?, quantity=? WHERE id=?;`,
			cartItem.Total, cartItem.Quantity, cartItem.ID).Exec()
	if err != nil {
		return err
	}

	return nil
}

func(c *CartItemRepository) Delete(id string) error {
	err := c.DB.Query(`Delete FROM cart_items WHERE id=?`,id).Exec()
	if err != nil {
		return err
	}

	return nil
}