package repository

import (
	"moonlay_technical_test/helper"
	"moonlay_technical_test/model"

	"github.com/stretchr/testify/mock"
)

type ListRepositoryMock struct {
	Mock mock.Mock
}

func (m *ListRepositoryMock) GetList(request helper.Pagination) (pagination helper.Pagination, err error) {
	arguments := m.Mock.Called()
	if arguments.Get(0) == nil {
		return pagination, err
	} else {
		list := arguments.Get(0).(helper.Pagination)
		return list, err
	}
}

func (m *ListRepositoryMock) GetDetailList(Id int) (list *model.List, err error) {
	arguments := m.Mock.Called(Id)
	if arguments.Get(0) == nil {
		return nil, err
	} else {
		list := arguments.Get(0).(model.List)
		return &list, nil
	}
}

func (m *ListRepositoryMock) GetListRelational() (lists []model.List, err error) {
	arguments := m.Mock.Called()
	if arguments.Get(0) == nil {
		return lists, err
	} else {
		list := arguments.Get(0).([]model.List)
		return list, err
	}
}

func (m *ListRepositoryMock) CreateList(list model.List) (err error) {
	arguments := m.Mock.Called(list)
	if arguments.Get(0) == nil {
		return err
	} else {
		return nil
	}
}

func (m *ListRepositoryMock) UpdateList(list model.List) (err error) {
	arguments := m.Mock.Called(list)
	if arguments.Get(0) == nil {
		return err
	} else {
		return nil
	}
}

func (m *ListRepositoryMock) DeleteList(Id int) (err error) {
	arguments := m.Mock.Called(Id)
	if arguments.Get(0) == nil {
		return err
	} else {
		return nil
	}
}
