package entitytest

import (
	"star_store/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product,err := entity.NewProduct("camisa","12314452","Juca","thum.url", 10)

	assert.Nil(t,err)
	assert.NotNil(t,product)
	assert.Equal(t,"camisa",product.Title)
	assert.Equal(t,"12314452",product.ZipCode)
	assert.Equal(t,"Juca",product.Seller)
	assert.Equal(t,"thum.url",product.ThumbnailHD)
	assert.Equal(t,10,product.Price)
}

func TestNewProductInvalidTitle(t *testing.T) {
	product,err := entity.NewProduct("","12314452","Juca","thum.url", 10)

	assert.Nil(t,product)
	assert.NotNil(t,err)
}

func TestNewProductInvalidZipCode(t *testing.T) {
	product,err := entity.NewProduct("title","","Juca","thum.url", 10)

	assert.Nil(t,product)
	assert.NotNil(t,err)
}

func TestNewProductInvalidSeller(t *testing.T) {
	product,err := entity.NewProduct("title","12314452","","thum.url", 10)

	assert.Nil(t,product)
	assert.NotNil(t,err)
}