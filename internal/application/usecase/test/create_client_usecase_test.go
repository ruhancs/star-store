package usecase_test

import (
	"star_store/internal/application/dto"
	"star_store/internal/application/usecase"
	mock_gateway "star_store/internal/application/usecase/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateClientUseCase(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	clientMockRepo := mock_gateway.NewMockClientRepositoryInterface(ctrl)
	input := dto.InputCreateClientDto{
		Name: "C1",
		ZipCode: "1348723",
	}
	clientMockRepo.EXPECT().Create(gomock.Any()).AnyTimes()

	createClientUseCase := usecase.NewCreateClientUseCase(clientMockRepo)

	out,err := createClientUseCase.Execute(input)

	assert.Nil(t,err)
	assert.NotNil(t,out)
	assert.Equal(t,"C1",out.Name)
	assert.Equal(t,"1348723",out.ZipCode)
}