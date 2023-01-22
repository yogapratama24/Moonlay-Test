package repository

import (
	"moonlay_technical_test/model"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) (*model.Category, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, nil
	} else {
		category := arguments.Get(0).(model.Category)
		return &category, nil
	}

}
