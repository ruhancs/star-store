package repository_test

import (
	"star_store/internal/domain/entity"
	"star_store/internal/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var clientRepository = repository.NewClientRepository(session)
var client,_ = entity.NewClient("N1","Z1")

func TestCreateClient(t *testing.T) {
	err := clientRepository.Create(client)

	assert.Nil(t,err)
	clientRepository.Delete(client.ID)
}

func TestGetClient(t *testing.T) {
	clientRepository.Create(client)
	client,err := clientRepository.Get(client.ID)

	assert.Nil(t,err)
	assert.Equal(t,"N1",client.Name)
	assert.Equal(t,"Z1",client.ZipCode)
	clientRepository.Delete(client.ID)
}