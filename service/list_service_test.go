package service

import (
	"moonlay_technical_test/model"
	"moonlay_technical_test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var listRepository = &repository.ListRepositoryMock{Mock: mock.Mock{}}
var listService = ListServices{ListRepository: listRepository}

func TestGetDetailListSuccess(t *testing.T) {
	list := model.List{
		Id:    1,
		Title: "Title 1",
	}
	listRepository.Mock.On("GetDetailList", 2).Return(list)
	result, err := listService.GetDetailList(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, list.Id, result.Id)
	assert.Equal(t, list.Title, result.Title)
}

func TestCreateListSuccess(t *testing.T) {
	listRequest := model.ListCreateRequest{
		Id:          1,
		Title:       "Title 1",
		Description: "description",
	}
	list := model.List{
		Id:          1,
		Title:       "Title 1",
		Description: "description",
	}
	listRepository.Mock.On("CreateList", list).Return(nil)
	err := listService.CreateList(listRequest)
	assert.NoError(t, err)
	assert.NotEmpty(t, listRequest)
	assert.Equal(t, listRequest.Title, list.Title)
	assert.Equal(t, listRequest.Id, list.Id)
	assert.Equal(t, listRequest.Title, list.Title)
}
func TestUpdateListSuccess(t *testing.T) {
	listRequest := model.ListUpdateRequest{
		Id:          1,
		Title:       "Title 1",
		Description: "description",
	}
	list := model.List{
		Id:          1,
		Title:       "Title 1",
		Description: "description",
	}
	listRepository.Mock.On("UpdateList", list).Return(nil)
	err := listService.UpdateList(listRequest)
	assert.NoError(t, err)
	assert.NotEmpty(t, listRequest)
	assert.Equal(t, listRequest.Title, list.Title)
	assert.Equal(t, listRequest.Id, list.Id)
	assert.Equal(t, listRequest.Title, list.Title)
}

func TestDeleteList(t *testing.T) {
	listRepository.Mock.On("DeleteList", 2).Return(nil)
	err := listService.DeleteList(2)
	assert.Nil(t, err)
}
