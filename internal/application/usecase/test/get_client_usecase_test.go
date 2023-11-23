package usecase_test

import (
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"star_store/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetClientUseCase(t *testing.T) {
	c, _ := entity.NewClient("C1", "287361")
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	clientMockRepo := mock_gateway.NewMockClientRepositoryInterface(ctrl)
	clientMockRepo.EXPECT().Get(c.ID).Return(c, nil)

	getClientUseCase := usecase.NewGetClientUseCase(clientMockRepo)

	out,err := getClientUseCase.Execute(c.ID)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,"C1",c.Name)
	assert.Equal(t,"287361",c.ZipCode)
}