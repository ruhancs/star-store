package entitytest

import (
	"star_store/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client,err := entity.NewClient("Ka","123325435")

	assert.Nil(t,err)
	assert.NotNil(t,client)
	assert.Equal(t,"Ka",client.Name)
	assert.Equal(t,"123325435",client.ZipCode)
}
