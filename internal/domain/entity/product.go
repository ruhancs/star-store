package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID          string  `json:"id" valid:"required"`
	Title       string  `json:"title" valid:"required"`
	Price       float32 `json:"price" valid:"required"`
	ZipCode     string  `json:"zip_code" valid:"required"`
	Seller      string  `json:"seller" valid:"required"`
	ThumbnailHD string  `json:"thumbnail" valid:"required"`
	Date        string  `json:"date"`
}

func NewProduct(title, zip, seller, thumb string, price float32) (*Product, error) {
	product := &Product{
		ID:          uuid.NewV4().String(),
		Title:       title,
		Price:       price,
		ZipCode:     zip,
		Seller:      seller,
		ThumbnailHD: thumb,
		Date:        time.Now().UTC().String(),
	}
	err := product.isValid()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}
