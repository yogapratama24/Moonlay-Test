package service

import (
	"moonlay_technical_test/repository"

	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// func TestCategoryNotFound(t *testing.T) {
// 	categoryRepository.Mock.On("FindById", "1").Return(nil)
// 	list, err := categoryService.Get("1")
// 	assert.Nil(t, list)
// 	assert.NotNil(t, err)
// }

// func TestGetCategorySuccess(t *testing.T) {
// 	list := model.Category{
// 		Id:   "1",
// 		Name: "Title 1",
// 	}
// 	categoryRepository.Mock.On("FindById", "1").Return(list)
// 	result, err := categoryService.Get("1")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, list.Id, result.Id)
// 	assert.Equal(t, list.Name, result.Name)
// }
