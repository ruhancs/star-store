package repository

import (
	"fmt"
	"star_store/internal/domain/entity"

	"github.com/gocql/gocql"
)

type ProductRepository struct {
	DB *gocql.Session
}

func NewProductRepository(db *gocql.Session) *ProductRepository{
	return &ProductRepository{
		DB: db,
	}
}

func(p *ProductRepository) Create(product *entity.Product) error {
	date := product.Date
	err := p.DB.Query(`
		INSERT INTO products(id, title, price, zip_code, seller, thumbnail, date) VALUES (?, ?, ?, ?, ?, ?, ?);`,
			product.ID,product.Title,product.Price,product.ZipCode,product.Seller,product.ThumbnailHD,date).Exec()
	if err != nil {
		return err
	}

	return nil
}

func(p *ProductRepository) List() ([]*entity.Product,error) {
	var products []*entity.Product
	rows := p.DB.Query("SELECT * from products;").Iter()
	defer rows.Close()
	scanner := rows.Scanner()
	for scanner.Next() {
		var product entity.Product
		err := scanner.Scan(&product.ID,&product.Date,&product.Price,&product.Seller,&product.ThumbnailHD,&product.Title,&product.ZipCode)
		if err != nil {
			return nil,err
		}
		products = append(products, &product)
	}

	fmt.Println(products)

	return products,nil
}

func(p *ProductRepository) Get(id string) (*entity.Product,error) {
	var product entity.Product
	err := p.DB.Query("SELECT * from products WHERE id=?;",id).Scan(&product.ID,&product.Date,&product.Price,&product.Seller,&product.ThumbnailHD,&product.Title,&product.ZipCode)
	if err != nil {
		return nil,err
	}

	return &product,nil
}

func(p *ProductRepository) Delete(id string) error {
	err := p.DB.Query("DELETE from products WHERE id=?;",id).Exec()
	if err != nil {
		return err
	}
	return nil
}