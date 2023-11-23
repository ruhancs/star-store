package repository_test

import (
	"star_store/internal/domain/entity"
	"star_store/internal/infra/db"
	"star_store/internal/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var session,_ = db.ConnetToCassandraCluster()
var productRepository = repository.NewProductRepository(session)
var product,_ = entity.NewProduct("P1","Z1","S1","url",20.0)
var product2,_ = entity.NewProduct("P2","Z2","S2","url",10.0)

func TestCreateProduct(t *testing.T) {
	err := productRepository.Create(product)

	assert.Nil(t,err)
	productRepository.Delete(product.ID)
}

func TestListProduct(t *testing.T) {
	productRepository.Create(product)
	productRepository.Create(product2)
	products,err := productRepository.List()

	assert.Nil(t,err)
	assert.NotNil(t,products)
	productRepository.Delete(product.ID)
	productRepository.Delete(product2.ID)
}

func TestGetProduct(t *testing.T) {
	productRepository.Create(product)
	product,err := productRepository.Get(product.ID)

	assert.Nil(t,err)
	assert.NotNil(t,product)
	assert.Equal(t,"P1",product.Title)
	assert.Equal(t,"Z1",product.ZipCode)
	assert.Equal(t,"S1",product.Seller)
	assert.Equal(t,"url",product.ThumbnailHD)
	assert.Equal(t,float32(20),product.Price)
}