package entitytest

import (
	"star_store/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidCreditCard(t *testing.T) {
	transaction,err := entity.NewTransaction("123","JJ","1235",10)

	assert.NotNil(t,err.Error())
	assert.Equal(t,"invalid credit card number",err.Error())
	assert.Nil(t,transaction)
}

func TestValidCreditCard(t *testing.T) {
	transaction,err := entity.NewTransaction("123","JJ","1234123412341234",10)

	assert.Nil(t,err)
	assert.NotNil(t,transaction)
}